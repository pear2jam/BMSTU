.MODEL SMALL
STACK 100h
.DATA
    ARR db 0, 1, 2, 4, 6, 7, 3,3,3,3,3,3    ; here s array
    len dw 12
    positive_info db 'Postive elements: $'
    negative_info db 'Negative elements: $'
.CODE
    mov ax, @Data
    mov ds, ax                      ; offseting data
    mov cx, [len]                   ; number of elements
    mov ah, 0                       ; count negative
    mov al, 0                       ; count positive
check_symbol:                       ; start loop body
    mov di, cx                      ; moving index in di because cx cant be index for []
    mov bl, byte ptr ARR[di-1]      ; movind element (from -128 to 127) in bl
    cmp bl, 0F0h                    ; if n > f0 -> n < 0
    jb add_positive                 ; command if bl > 240
    inc ah                          ; increment negative
    loop check_symbol               ; back to loop body
add_positive:                       ; if bl < 240
    inc al                          ; increment positive
    loop check_symbol               ; back to loop body

    mov bx, 10                      ; for get_decimal

    push ax                         ; negative count on stack

    mov dx, offset positive_info
    mov ah, 09h
    int 21h
    pop ax
    call get_decimal

    mov ah, 4ch ; ending program
    int 21h

get_decimal PROC             ; add to stack decimal form of ax
    xor dx, dx
    div bx
    add dl, '0'
    push dx
    inc cx
    test ax, ax
    jnz get_decimal
print:
    pop dx
    mov ah, 02h
    int 21h
    loop print
    ret
get_decimal ENDP

    END
