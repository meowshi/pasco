const RusToEn = new Map([
  ["Ё", "`"],
  ["Й", "Q"],
  ["Ц", "W"],
  ["У", "E"],
  ["К", "R"],
  ["Е", "T"],
  ["Н", "Y"],
  ["Г", "U"],
  ["Ш", "I"],
  ["Щ", "O"],
  ["З", "P"],
  ["Х", "["],
  ["Ъ", "]"],
  ["Ф", "A"],
  ["Ы", "S"],
  ["В", "D"],
  ["А", "F"],
  ["П", "G"],
  ["Р", "H"],
  ["О", "J"],
  ["Л", "K"],
  ["Д", "L"],
  ["Ж", ";"],
  ["Э", "'"],
  ["Я", "Z"],
  ["Ч", "X"],
  ["С", "C"],
  ["М", "V"],
  ["И", "B"],
  ["Т", "N"],
  ["Ь", "M"],
  ["Б", ","],
  ["Ю", "."],
]);

export function rfidToKey(rfid) {
  rfid = rfid.toUpperCase();
  let result = "";
  for (let i = 0; i < rfid.length; i++) {
    if (rfid[i] >= "0" && rfid[i] <= "9") {
      result += rfid.charAt(i);
    } else {
      const eng = RusToEn.get(rfid[i]);
      if (eng === undefined) {
        result += rfid.charAt(i);
      } else {
        result += eng;
      }
    }
  }
  return result;
}
