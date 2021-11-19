### MIME library의 Decode함수 분석

```
word := `=?utf-8?q?=C2=A1Hola,_se=C3=B1or!?=`

1.   utf-8?q?=C2=A1Hola,_se=C3=B1or!
2.   5
3.   utf-8
4.         q
5.           =C2=A1Hola,_se=C3=B1or!
6. ¡Hola, señor!
```
