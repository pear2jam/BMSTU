.MODEL SMALL
STACK 100h
.DATA
    ARR db 0, 1, -2, 4, -6, 7 ; here s array
.CODE
    mov ax, @Data
    mov ds, ax ; offseting data
    mov cx, 6h ; number of elements
    mov ah, 0 ; count negative
    mov al, 0 ; count positive
check_symbol: ; start loop body
    mov di, cx ; moving index in di because cx cant be index for []
    mov bl, byte ptr ARR[di-1] ; movind element (from -16 to 15) in bl
    cmp bl, 240 ; 240=fo  if n > f0 -> n < 0
    jb add_positive ; command if bl > 240
    inc ah ; increment negative
    loop check_symbol ; back to loop body
add_positive: ; if bl < 240
    inc al ; increment positive
    loop check_symbol ; back to loop body

    mov ah, 4ch ; ending program
    int 21h
    END
