import glfw
from OpenGL.GL import *
from random import random
import time
import numpy as np

ang = 45
mode = 0  # from 0 to 2

vertices = [
    [-1, 1, -1],
    [1, 1, -1],
    [1, -1, -1],
    [-1, -1, -1],
    [-1, 1, 1],
    [1, 1, 1],
    [1, -1, 1],
    [-1, -1, 1]
]

polygons = [
    [0, 1, 2, 3],
    [1, 0, 4, 5],
    [1, 5, 2, 6],
    [7, 6, 2, 3],
    [7, 6, 5, 4],
    [3, 7, 4, 0]

]

edges = [
    [0, 1],
    [1, 5],
    [4, 5],
    [0, 4],
    [1, 2],
    [5, 6],
    [4, 7],
    [0, 3],
    [3, 7],
    [2, 3],
    [2, 6],
    [6, 7]
]

colors = [
    [1, 0, 0, 1],
    [0, 1, 0, 1],
    [0, 0, 1, 1],
    [0.6, 0.6, 0.6, 1],
    [0.7, 0.3, 0.8, 1],
    [1, 0.8, 0.2, 1]
]

mat_front = np.eye(4)
mat_front[2][2] = -1

mat_side = np.zeros((4, 4))
mat_side[2][0] = -1
mat_side[1][1] = 1
mat_side[0][2] = -1
mat_side[3][3] = 1

mat_top = np.zeros((4, 4))
mat_top[0][0] = 1
mat_top[2][1] = -1
mat_top[1][2] = -1
mat_top[3][3] = 1


def main():
    if not glfw.init():
        return
    window = glfw.create_window(640, 640, "Lab1", None, None)
    if not window:
        glfw.terminate()
        return
    glfw.make_context_current(window)
    #glfw.set_key_callback(window, key_callback)
    while not glfw.window_should_close(window):
        display(window)
    glfw.destroy_window(window)
    glfw.terminate()

def display(window):
    #time.sleep(0.02)
    global ang
    glClear(GL_COLOR_BUFFER_BIT)
    glLoadIdentity()
    glScaled(0.5, 0.5, 0.5)
    glRotated(ang, 1, 1, 1)
    #ang += 1
    glClearColor(1.0, 1.0, 1.0, 1.0)
    glPushMatrix()
    glRotatef(ang, 0, 0, 1)
    glBegin(GL_QUADS)
    # glColor3f(*colors[0])
    for ind, poly in enumerate(polygons):
        v_1 = vertices[poly[0]]
        v_2 = vertices[poly[1]]
        v_3 = vertices[poly[2]]
        v_4 = vertices[poly[3]]

        glColor4f(*colors[ind])
        glVertex3f(*v_1)

        glColor4f(*colors[ind])
        glVertex3f(*v_2)

        glColor4f(*colors[ind])
        glVertex3f(*v_3)

        glColor4f(*colors[ind])
        glVertex3f(*v_4)
    glEnd()

    glBegin(GL_LINES)
    for edge in edges:
        v_1 = vertices[edge[0]]
        v_2 = vertices[edge[1]]

        glVertex3f(*v_1)
        glColor3f(0.5, 0.5, 0.5)
        glVertex3f(*v_2)
        glColor3f(0.5, 0.5, 0.5)
    glEnd()
    glPopMatrix()

    glMatrixMode(GL_PROJECTION)
    glLoadMatrixf(mat_side)

    glMatrixMode(GL_MODELVIEW)
    glfw.swap_buffers(window)
    glfw.poll_events()


main()
