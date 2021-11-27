.MODEL SMALL
STACK 100h
.DATA
    A DW 4          ; (a + b/c)*d - 4
    B DW 3
    C DW 2
    D DW 3
    info_dec db 'In Dec: $'
    info_hex db 10, 'In Hex: $'
.CODE
    mov ax, @Data
    mov ds, ax
    mov ax, [B]
    mov dx, 0
    div [C]
    add ax, [A]
    mul [D]
    sub ax, 4

    mov bx, 10

    push ax
    push ax

    mov ah, 09h
    mov dx, offset info_dec
    int 21h

    pop ax
    call get_decimal

    mov ah, 09h
    mov dx, offset info_hex
    int 21h

    mov bx, 010h
    pop ax
    call get_decimal

    mov ah, 4ch
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
