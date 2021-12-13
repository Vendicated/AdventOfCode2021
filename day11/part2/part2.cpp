#include <cstdio>
#include <fstream>
#include <sstream>
#include <unordered_set>
#include <vector>

typedef std::vector<std::vector<int>> MapType;
void doField(MapType &map, int i, int j,
             std::vector<std::vector<int>> &flashedFields);

int main(void) {
  std::ifstream infile("day11/input.txt");
  MapType map;
  std::string line;

  while (std::getline(infile, line)) {
    std::vector<int> row;
    for (const auto &c : line) {
      row.push_back(c - '0');
    }
    map.push_back(row);
  }
  infile.close();

  int result;
  for (;;) {
    std::vector<std::vector<int>> flashedFields;
    for (int i = 0, size = map.size(); i < size; i++) {
      auto row = map[i];
      for (int j = 0, rowSize = row.size(); j < rowSize; j++) {
        doField(map, i, j, flashedFields);
      }
    }

    result++;
    if (flashedFields.size() == map.size() * map[0].size()) {
      break;
    }
    for (const auto &field : flashedFields) {
      map[field[0]][field[1]] = 0;
    }
  }

  printf("%d", result);
  return 0;
}

void doField(MapType &map, int i, int j,
             std::vector<std::vector<int>> &flashedFields) {
  if (++map[i][j] > 9) {
    for (const auto &field : flashedFields) {
      if (field[0] == i && field[1] == j) {
        return;
      }
    }
    flashedFields.push_back({i, j});
    for (int m = -1; m < 2; m++) {
      for (int n = -1; n < 2; n++) {
        if (m == 0 && n == 0) {
          continue;
        }
        int i2 = i + m;
        int j2 = j + n;
        if (i2 < 0 || j2 < 0 || i2 > map.size() - 1 || j2 > map[0].size() - 1) {
          continue;
        }
        doField(map, i2, j2, flashedFields);
      }
    }
  }
}
