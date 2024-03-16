import glfw
from OpenGL.GL import *
from OpenGL.GLU import *
import numpy as np
from scipy import signal
import math


def key_callback(window, key, scancode, action, mods):
    if action == glfw.PRESS:
        if key == glfw.KEY_SPACE:
            drawer.clear_points()
        if key == glfw.KEY_F:
            drawer.fill = not drawer.fill
        if key == glfw.KEY_L:
            drawer.use_post_filter = not drawer.use_post_filter


def mouse_button_callback(window, button, action, mods):
    if button == glfw.MOUSE_BUTTON_LEFT and action == glfw.PRESS:
        pos = glfw.get_cursor_pos(window)
        print(pos[0])
        print(pos[1])
        coords = [pos[0] , height//2 - pos[1]]
        drawer.add_point(coords)
        drawer.points_pixels = [[*relative_to_pixels(*p)] for p in drawer.points]


def relative_to_pixels(x, y):
    return int(x), int(y)


def conv2d(a, f):
    s = f.shape + tuple(np.subtract(a.shape, f.shape) + 1)
    strd = np.lib.stride_tricks.as_strided
    subM = strd(a, shape=s, strides=a.strides * 2)
    return np.einsum('ij,ijkl->kl', f, subM)


class Drawer:
    def __init__(self):
        self.points = []
        self.points_pixels = [[*relative_to_pixels(*p)] for p in self.points]
        self.pixels = np.ones([width, height, 3], dtype=GLfloat)
        self.peregorodka_x = width // 2
        self.fill = False
        self.use_post_filter = False
        self.prev_point = []
        self.next_point = []
        print(self.pixels.shape)

    def display(self, window):
        glClearColor(1.0, 1.0, 1.0, 1.0)
        glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT)
        glMatrixMode(GL_PROJECTION)
        glLoadIdentity()
        gluOrtho2D(0.0, 750, 0.0, 750)
        glMatrixMode(GL_MODELVIEW)
        glLoadIdentity()

        self.pixels = np.ones([width, height, 3], dtype=GLfloat)
        self.draw_edges()
        if self.fill:
            self.fill_polygon()
        if self.use_post_filter:
            self.post_filter()

        pixels_vector = np.ascontiguousarray(self.pixels).reshape(width * height * 3)
        glDrawPixels(width, height, GL_RGB, GL_FLOAT, (GLfloat * len(pixels_vector))(*pixels_vector))

        glfw.swap_buffers(window)
        glfw.poll_events()

    def draw_polygon_edges(self):
        glLineWidth(2)
        glColor3f(1, 0, 0)
        glBegin(GL_POLYGON)
        for coord in self.points:
            glVertex2f(*coord)
        glEnd()

    def fill_edge(self, x0, y0, x1, y1):
        if y0 == y1 or x0 == x1:
            return None
        if y0 > y1:
            x0, x1, y0, y1 = x1, x0, y1, y0
        line = lambda y: ((y - y0) / (y1 -y0)) * (x1 - x0) + x0
        for y in range(y0, y1):
            x = math.ceil(line(y))
            self.pixels[y, x, :] = [1, 0, 0]

    def draw_edges(self):
        for i in range(len(self.points)):
            if i == len(self.points) - 1:
                x0, y0 = relative_to_pixels(*self.points[i])
                x1, y1 = relative_to_pixels(*self.points[0])
                self.fill_edge(x0, y0, x1, y1)
            else:
                x0, y0 = relative_to_pixels(*self.points[i])
                x1, y1 = relative_to_pixels(*self.points[i + 1])
                self.fill_edge(x0, y0, x1, y1)

    def fill_polygon(self):
        print(self.points_pixels)
        for y in range(0, height):
            inside = False
            for x in range(0, width):
                skip = False
                if[x, y] in self.points_pixels:
                    ind = self.points_pixels.index([x, y])
                    prev_ind = ind - 1
                    next_ind = ind + 1
                    if ind == 0:
                        prev_ind = -1
                    if ind == len(self.points_pixels)-1:
                        next_ind = 0
                    print(x, y)
                    if (self.points_pixels[prev_ind][1]-y)*(self.points_pixels[next_ind][1]-y) > 0:
                        skip = True
                        print('skip')
                if (self.pixels[y][x] == [1, 0, 0]).all():
                    if not skip:
                        inside = not inside
                if inside:
                    self.pixels[y][x] = [1, 0, 0]

    def post_filter(self):
        w_k = (1 / 6) * np.array([[0, 1, 0],
                                  [1, 2, 1],
                                  [0, 1, 0], ],
                                 dtype='float')
        for i in range(self.pixels.shape[2]):
            self.pixels[:, :, i] = signal.convolve2d(np.pad(self.pixels[:, :, i], 1), w_k, 'valid')

    def add_point(self, point):
        self.points.append(point)

    def clear_points(self):
        self.points = []


if __name__ == '__main__':
    if not glfw.init():
        exit()
    width, height = 1500, 1500
    window = glfw.create_window(width // 2, height // 2, "lab4", None, None)
    if not window:
        glfw.terminate()
        exit()
    glfw.make_context_current(window)
    glfw.set_input_mode(window, glfw.STICKY_KEYS, GL_TRUE)
    glfw.set_key_callback(window, key_callback)
    glfw.set_mouse_button_callback(window, mouse_button_callback)
    glPolygonMode(GL_FRONT_AND_BACK, GL_LINE)

    drawer = Drawer()
    glViewport(0, 0, width, height)

    while not glfw.window_should_close(window):
        drawer.display(window)
    glfw.destroy_window(window)
    glfw.terminate()
