import 'dart:convert';
import 'dart:io';

void main(List<String> arguments) async {
  var input = await readFile("24-test.txt");

  print(input);

  /*
  https://github.com/mrphlip/aoc/blob/master/2021/24.md

  z.push(A+5)
  z.push(B+14)
  z.push(C+15)
  z.push(D+16)

  // none of the conditional stack pushes must be allowed to happen
  if E != z.pop() - 16 ? z.push(E+8)
  if F != z.pop() - 11 ? z.push(F+9)
  if G != z.pop() - 6 ? z.push(G+2)
  z.push(H+13)
  z.push(I+16)
  if J != z.pop() - 10 ? z.push(J+6)
  if K != z.pop() - 8 ? z.push(K+6)
  if L != z.pop() - 11 ? z.push(L+9)
  z.push(M+11)
  if N!= z.pop() - 15 ? z.push(N+5)

  E == D
  F == C + 4
  G == B + 8
  J == I + 6
  K == H + 5
  L == A - 6
  N == M - 4

  // largest
  ABCDEFGHIJKLMN
  91599994399395

  // smallest
  ABCDEFGHIJKLMN
  71111591176151
  */
}

Future<List<String>> readFile(String path) async {
  final lines =
      utf8.decoder.bind(File(path).openRead()).transform(const LineSplitter());

  return await lines.toList();
}
