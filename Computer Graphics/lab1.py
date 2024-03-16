import glfw
from OpenGL.GL import *
from random import random
import time

triangles = [(0, 0), (0.5, 0.5), (-0.7, 0.3), (0.7, -0.3)]
colors_1 = [(0.1, 0.1, 0.1), (0.35, 0.0, 0.89), (0.0, 1.0, 1.0)]
colors_2 = [(0.5, 0.4, 0.2), (0.3, 0.0, 0.1), (0.5, 0, 0)]

size = 0.1
shift_x = size
shift_y = size
max_len = 12
mode = 1


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

def display(window):
    global triangles
    global shift_x, shift_y
    glClear(GL_COLOR_BUFFER_BIT)
    glLoadIdentity()
    glClearColor(1.0, 1.0, 1.0, 1.0)
    glPushMatrix()
    #glRotatef(angle, 0, 0, 1)
    glBegin(GL_TRIANGLES)

    if mode:
        colors = colors_1
    else:
        colors = colors_2


    for i in range(len(triangles) - 2):
        glColor3f(*colors[0])
        glVertex2f(triangles[i][0], triangles[i][1])
        glColor3f(*colors[1])
        glVertex2f(triangles[i+1][0], triangles[i+1][1])
        glColor3f(*colors[2])
        glVertex2f(triangles[i+2][0], triangles[i+2][1])
    glEnd()

    if mode:
        to_x = triangles[-1][0] + shift_x + random()*size*7
        to_y = triangles[-1][1] + shift_y + random()*size*7
    else:
        to_x = triangles[0][0] + shift_x + random()*size*7
        to_y = triangles[0][1] + shift_y + random()*size*7

    shift_x += (random()*size - size/2)
    shift_y += (random()*size - size/2)

    if 1 - to_x < 0.2:
        shift_x -= random()*size*1.5
    if to_x - (-1) < 0.2:
        shift_x += random()*size*1.5

    if 1 - to_y < 0.3:
        shift_y -= random()*size*1.5
    if to_y - (-1) < 0.3:
        shift_y += random()*size*1.5

    if to_x > 1:
        if mode:
            to_x = triangles[-1][0] - random()*size
        else:
            to_x = triangles[0][0] - random()*size
    if to_x < -1:
        if mode:
            to_x = triangles[-1][0] + random()*size
        else:
            to_x = triangles[0][0] + random()*size

    if to_y > 1:
        if mode:
            to_y = triangles[-1][1] - random()*size
        else:
            to_y = triangles[0][1] - random() * size
    if to_y < -1:
        if mode:
            to_y = triangles[-1][1] + random()*size
        else:
            to_y = triangles[0][1] + random() * size

    if mode:
        triangles.append((to_x, to_y))
    else:
        triangles.insert(0, (to_x, to_y))

    #triangles.append((random()*size - size/2, random()*size - size/2))
    if len(triangles) > max_len:
        if mode:
            triangles = triangles[1:]
        else:
            triangles = triangles[:-1]

    time.sleep(0.1)


    glPopMatrix()
    glfw.swap_buffers(window)
    glfw.poll_events()


def key_callback(window, key, scancode, action,  mods):
    global mode
    if action == glfw.PRESS:
        if key == 32: # backspace
            mode = 1 - mode


main()
