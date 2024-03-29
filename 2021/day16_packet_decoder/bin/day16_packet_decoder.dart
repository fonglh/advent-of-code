import 'dart:math';

void main(List<String> arguments) {
  final rawInput =
      "005473C9244483004B001F79A9CE75FF9065446725685F1223600542661B7A9F4D001428C01D8C30C61210021F0663043A20042616C75868800BAC9CB59F4BC3A40232680220008542D89B114401886F1EA2DCF16CFE3BE6281060104B00C9994B83C13200AD3C0169B85FA7D3BE0A91356004824A32E6C94803A1D005E6701B2B49D76A1257EC7310C2015E7C0151006E0843F8D000086C4284910A47518CF7DD04380553C2F2D4BFEE67350DE2C9331FEFAFAD24CB282004F328C73F4E8B49C34AF094802B2B004E76762F9D9D8BA500653EEA4016CD802126B72D8F004C5F9975200C924B5065C00686467E58919F960C017F00466BB3B6B4B135D9DB5A5A93C2210050B32A9400A9497D524BEA660084EEA8EF600849E21EFB7C9F07E5C34C014C009067794BCC527794BCC424F12A67DCBC905C01B97BF8DE5ED9F7C865A4051F50024F9B9EAFA93ECE1A49A2C2E20128E4CA30037100042612C6F8B600084C1C8850BC400B8DAA01547197D6370BC8422C4A72051291E2A0803B0E2094D4BB5FDBEF6A0094F3CCC9A0002FD38E1350E7500C01A1006E3CC24884200C46389312C401F8551C63D4CC9D08035293FD6FCAFF1468B0056780A45D0C01498FBED0039925B82CCDCA7F4E20021A692CC012B00440010B8691761E0002190E21244C98EE0B0C0139297660B401A80002150E20A43C1006A0E44582A400C04A81CD994B9A1004BB1625D0648CE440E49DC402D8612BB6C9F5E97A5AC193F589A100505800ABCF5205138BD2EB527EA130008611167331AEA9B8BDCC4752B78165B39DAA1004C906740139EB0148D3CEC80662B801E60041015EE6006801364E007B801C003F1A801880350100BEC002A3000920E0079801CA00500046A800C0A001A73DFE9830059D29B5E8A51865777DCA1A2820040E4C7A49F88028B9F92DF80292E592B6B840";
  final binaryString = toBinary(rawInput);
  // print(binaryString);

  final testInput = "9C005AC2F8F0";
  final testBinaryString = toBinary(testInput);
  // print(testBinaryString);

  // Part 1
  print(sumVersionNumbers(binaryString).x.toInt());

  // Part 2
  print(
      "--------------------------------------------------------------------------------");
  print(evaluatePacket(binaryString).x.toInt());
}

// Cannot use built-in radix conversions as the padding gets lost
String toBinary(String hexString) {
  String result = "";

  Map<String, String> hexToBin = {
    "0": "0000",
    "1": "0001",
    "2": "0010",
    "3": "0011",
    "4": "0100",
    "5": "0101",
    "6": "0110",
    "7": "0111",
    "8": "1000",
    "9": "1001",
    "A": "1010",
    "B": "1011",
    "C": "1100",
    "D": "1101",
    "E": "1110",
    "F": "1111",
  };

  hexString.split("").forEach((element) {
    result += hexToBin[element]!;
  });
  return result;
}

// For part 2, parse packet and evaluate the expression from the literal values
// Use Point as a tuple, returning (literalResult, processedLength)
Point evaluatePacket(String packetBits) {
  // Insufficient bits to form the shortest literal packet
  if (packetBits.length < 11) {
    return Point(0, packetBits.length);
  }

  int processedLength = 0;
  List<int> literalValues = [];

  // Read packet header and drop those 6 bits
  int packetVersion = int.parse(packetBits.substring(0, 3), radix: 2);
  int packetTypeId = int.parse(packetBits.substring(3, 6), radix: 2);
  packetBits = packetBits.substring(6);
  processedLength += 6;

  switch (packetTypeId) {
    // Literal value packet
    case 4:
      String literalBits = "";
      while (true) {
        bool stopParsing = packetBits[0] != "1";
        literalBits += packetBits.substring(0, 5).substring(1);

        // Remove each group of 5 bits
        packetBits = packetBits.substring(5);
        processedLength += 5;
        if (stopParsing) {
          break;
        }
      }
      int literalValue = int.parse(literalBits, radix: 2);
      // print(literalValue);
      return Point(literalValue, processedLength);
    // Operator packet
    default:
      // Length type ID is the next bit
      String lengthTypeId = packetBits[0];
      packetBits = packetBits.substring(1);
      processedLength += 1;

      if (lengthTypeId == "0") {
        int subPacketLength = int.parse(packetBits.substring(0, 15), radix: 2);
        packetBits = packetBits.substring(15);
        processedLength += 15;

        int subPacketProcessedLength = 0;
        while (subPacketProcessedLength < subPacketLength) {
          var parseResult = evaluatePacket(packetBits);
          literalValues.add(parseResult.x.toInt());
          processedLength += parseResult.y.toInt();
          subPacketProcessedLength += parseResult.y.toInt();
          packetBits = packetBits.substring(parseResult.y.toInt());
        }
      } else {
        int subPacketCount = int.parse(packetBits.substring(0, 11), radix: 2);
        packetBits = packetBits.substring(11);
        processedLength += 11;

        int subPacketProcessedCount = 0;
        while (subPacketProcessedCount < subPacketCount) {
          var parseResult = evaluatePacket(packetBits);
          literalValues.add(parseResult.x.toInt());
          processedLength += parseResult.y.toInt();
          subPacketProcessedCount++;
          packetBits = packetBits.substring(parseResult.y.toInt());
        }
      }
      break;
  }

  // Process collected literal values according to operator packet type
  switch (packetTypeId) {
    // Sum packet
    case 0:
      return Point(literalValues.reduce((value, element) => value + element),
          processedLength);
    // Product packet
    case 1:
      return Point(literalValues.reduce((value, element) => value * element),
          processedLength);
    // Min packet
    case 2:
      return Point(literalValues.reduce(min), processedLength);
    // Max packet
    case 3:
      return Point(literalValues.reduce(max), processedLength);
    // Greater than
    case 5:
      return Point(
          literalValues[0] > literalValues[1] ? 1 : 0, processedLength);
    // Less than
    case 6:
      return Point(
          literalValues[0] < literalValues[1] ? 1 : 0, processedLength);
    // Equal
    case 7:
      return Point(
          literalValues[0] == literalValues[1] ? 1 : 0, processedLength);
    default:
      return Point(0, processedLength);
  }
}

// For part 1, parse packet and sum up the version numbers
// Use Point as a tuple, returning (versionNumberSum, processedLength)
Point sumVersionNumbers(String packetBits) {
  // Insufficient bits to form the shortest literal packet
  if (packetBits.length < 11) {
    return Point(0, packetBits.length);
  }

  int processedLength = 0;

  // Read packet header and drop those 6 bits
  int packetVersion = int.parse(packetBits.substring(0, 3), radix: 2);
  int packetTypeId = int.parse(packetBits.substring(3, 6), radix: 2);
  packetBits = packetBits.substring(6);
  processedLength += 6;

  switch (packetTypeId) {
    // Literal value packet
    case 4:
      String literalBits = "";
      while (true) {
        bool stopParsing = packetBits[0] != "1";
        literalBits += packetBits.substring(0, 5).substring(1);

        // Remove each group of 5 bits
        packetBits = packetBits.substring(5);
        processedLength += 5;
        if (stopParsing) {
          break;
        }
      }
      int literalValue = int.parse(literalBits, radix: 2);
      // print(literalValue);
      return Point(packetVersion, processedLength);
    // Operator packet
    default:
      // Length type ID is the next bit
      String lengthTypeId = packetBits[0];
      packetBits = packetBits.substring(1);
      processedLength += 1;

      if (lengthTypeId == "0") {
        int subPacketLength = int.parse(packetBits.substring(0, 15), radix: 2);
        packetBits = packetBits.substring(15);
        processedLength += 15;

        int subPacketProcessedLength = 0;
        while (subPacketProcessedLength < subPacketLength) {
          var parseResult = sumVersionNumbers(packetBits);
          packetVersion += parseResult.x.toInt();
          processedLength += parseResult.y.toInt();
          subPacketProcessedLength += parseResult.y.toInt();
          packetBits = packetBits.substring(parseResult.y.toInt());
        }
      } else {
        int subPacketCount = int.parse(packetBits.substring(0, 11), radix: 2);
        packetBits = packetBits.substring(11);
        processedLength += 11;

        int subPacketProcessedCount = 0;
        while (subPacketProcessedCount < subPacketCount) {
          var parseResult = sumVersionNumbers(packetBits);
          packetVersion += parseResult.x.toInt();
          processedLength += parseResult.y.toInt();
          subPacketProcessedCount++;
          packetBits = packetBits.substring(parseResult.y.toInt());
        }
      }
      break;
  }
  return Point(packetVersion, processedLength);
}
