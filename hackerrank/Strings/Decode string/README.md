## Decode string

Dada uma string codificada, retorne a string decodificada.

A regra de codificação é: k[string_codificada], onde a string_codificada dentro dos colchetes serão repetidas o número de k vezes. O valor de k será sempre um número positivo.
Você deve assumir que as strings de entrada são sempre válidas, sem espaço e os colchetes estão bem formatados.

### Exemplos:

* s = "2[a]3[bc]", retornará "aabcbcbc".
* s = "3[a2[c]]", retornará "accaccacc".
* s = "2[abc]3[cd]ef", retornará "abcabccdcdcdef".