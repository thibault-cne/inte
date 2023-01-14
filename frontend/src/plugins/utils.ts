export function splitIntoDigits(num: number): string[] {
  // convert num i nto float with 2 decimals
  const parse_num = parseInt(num.toString());
  const digits = parse_num.toString().split("");
  // Add a . every 3 digits (ignoring the decimals)
  for (let i = digits.length - 3; i > 0; i -= 3) {
    digits.splice(i, 0, " ");
  }
  return digits;
}
