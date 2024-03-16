import glfw
from OpenGL.GL import *
from random import random
import time
import numpy as np

ang_x = 0
ang_y = 0
ang_z = 0

mode = 0  # from 0 to 2

vertices = [
    [-1, 1, -1],
    [1, 1, -1], # 1
    [1, -1, -1], # 2
    [-1, -1, -1],
    [-1, 1, 1],
    [1, 1, 1], # 5
    [1, -1, 1], # 6
    [-1, -1, 1]
]

polygons = [
    [0, 1, 2, 3],
    [1, 0, 4, 5],
    [1, 2, 6, 5],
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
    [1, 0, 0],
    [0, 1, 0],
    [0, 0, 1],
    [0.6, 0.6, 0.6],
    [0.7, 0.3, 0.8],
    [1, 0.8, 0.2]
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
    glfw.set_key_callback(window, key_callback)
    while not glfw.window_should_close(window):
        display(window)
    glfw.destroy_window(window)
    glfw.terminate()

def draw_cube(trans_vec):
    global ang_x, ang_y, ang_z


    glPushMatrix()

    glLoadIdentity()
    glTranslatef(*trans_vec)
    glClearColor(1.0, 1.0, 1.0, 1.0)

    glRotatef(ang_x, 1, 0, 0)
    glRotatef(ang_y, 0, 1, 0)
    glRotatef(ang_z, 0, 0, 1)

    glScaled(0.3, 0.3, 0.3)
    glBegin(GL_QUADS)
    for ind, poly in enumerate(polygons):

        v_1 = vertices[poly[0]]
        v_2 = vertices[poly[1]]
        v_3 = vertices[poly[2]]
        v_4 = vertices[poly[3]]

        glColor3f(*colors[ind])
        glVertex3f(*v_1)

        glColor3f(*colors[ind])
        glVertex3f(*v_2)

        glColor3f(*colors[ind])
        glVertex3f(*v_3)

        glColor3f(*colors[ind])
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


def display(window):
    glClear(GL_COLOR_BUFFER_BIT)

    glMatrixMode(GL_PROJECTION)
    glLoadMatrixf(mat_front)
    glMatrixMode(GL_MODELVIEW)
    draw_cube([-0.5,0.5,0])
    glMatrixMode(GL_PROJECTION)
    glLoadMatrixf(mat_top)
    glMatrixMode(GL_MODELVIEW)
    draw_cube([0.5, 0, -0.5])
    glMatrixMode(GL_PROJECTION)
    glLoadMatrixf(mat_side)
    glMatrixMode(GL_MODELVIEW)
    draw_cube([0, -0.5, -0.5])

    #glPopMatrix()

    #glMatrixMode(GL_MODELVIEW)
    glfw.swap_buffers(window)
    glfw.poll_events()


def key_callback(window, key, scancode, action,  mods):
    global ang_x, ang_y, ang_z
    if action == glfw.REPEAT:
        if key == 81: # q
            ang_x += 2
        if key == 87: # w
            ang_y += 2
        if key == 69: # e
            ang_z += 2



main()
