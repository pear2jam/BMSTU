.MODEL SMALL
.STACK 100h
.DATA
    A DW 4h ; (a + b/c)*d - 4
    B DW 3h
    C DW 2h
    D DW 3h
.CODE
    mov ax, @Data
    mov ds, ax
    mov ax, [B]
    mov dx, 0 ;чтобы не обнулить старшие биты делимого
    div [C]
    add ax, [A]
    mul [D]
    sub ax, 4
    mov dx, ax ;здесь хранится результат

    mov ah, 4ch
    int 21h
    END
