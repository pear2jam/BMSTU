.MODEL SMALL
STACK 100h
.DATA
    ARR db 0, -1, 2, 4, -6, 7, 3, 3    ; here s array
    len dw 8
    positive_info db 'Postive elements: $'
    negative_info db 10, 'Negative elements: $'
.CODE
    mov ax, @Data
    mov ds, ax                      
    mov cx, [len]                   
    mov ah, 0                       ; count negative
    mov al, 0                       ; count positive
check_symbol:
    mov di, cx                      ; moving index in di because cx cant be index for []
    mov bl, byte ptr ARR[di-1]      ; movind element (from -128 to 127) in bl
    cmp bl, 0F0h                    ; if n > f0 -> n < 0
    jb add_positive                 ; if bl > fo
    inc ah
    loop check_symbol
add_positive:                       ; if bl < f0
    inc al
    loop check_symbol

    mov bx, 10                      ; for get_decimal

    push ax                         ; negative count on stack (negative call)
    push ax                         ; positive call
    
    mov dx, offset positive_info
    mov ah, 09h
    int 21h

    pop ax
    xor ah, ah
    call get_decimal

    mov dx, offset negative_info
    mov ah, 09h
    int 21h

    pop ax
    mov al, ah
    xor ah, ah
    call get_decimal

    mov ah, 4ch ; ending program
    int 21h

get_decimal PROC                    ; add to stack decimal form of ax
    xor dx, dx
    div bx
    add dl, '0'
    push dx
    inc cx
    test ax, ax
    jnz get_decimal
print:                              ; print stack
    pop dx
    mov ah, 02h
    int 21h
    loop print
    ret
get_decimal ENDP

    END
