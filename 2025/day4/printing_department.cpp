#include <iostream>
#include <fstream>
#include <sstream>

constexpr int dirs[8][2] = {
    {-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}};

int64_t get_accessible_rolls(std::vector<std::vector<char>> &wall,
                             size_t rows, size_t cols)
{
    int64_t total_accessible_rolls{0};
    std::vector<std::pair<int, int>> markers;

    for (int32_t i = 0; i < rows; i++)
    {
        for (int32_t j = 0; j < cols; j++)
        {
            if (wall[i][j] != '.')
            {
                int32_t rolls{0};

                for (const auto &dir : dirs)
                {
                    int32_t nr = i + dir[0];
                    int32_t nc = j + dir[1];

                    if (nr >= 0 && nr < rows && nc >= 0 && nc < cols && wall[nr][nc] == '@')
                        rolls++;
                }

                if (rolls < 4)
                {
                    markers.emplace_back(i, j);
                    total_accessible_rolls++;
                }
            }
        }
    }

    for (const auto &pair : markers)
    {
        wall[pair.first][pair.second] = '.';
    }

    return total_accessible_rolls;
}

int main()
{
    std::ifstream file("input.txt");

    if (!file)
    {
        std::cout << "Error opening input file!" << std::endl;
        return 1;
    }

    std::string line;
    std::vector<std::vector<char>> wall;

    while (std::getline(file, line))
    {
        wall.emplace_back(line.begin(), line.end());
    }

    size_t rows{wall.size()}, cols{wall[0].size()};

    int64_t current_accessible_rolls{get_accessible_rolls(wall, rows, cols)};

    std::cout << "(PART 1) Total accessible rolls by forklift : "
              << current_accessible_rolls << std::endl;

    int64_t total_accessible_rolls{0};

    while (current_accessible_rolls > 0)
    {
        total_accessible_rolls += current_accessible_rolls;
        current_accessible_rolls = get_accessible_rolls(wall, rows, cols);
    }

    std::cout << "(PART 2) Total accessible rolls by forklift : "
              << total_accessible_rolls << std::endl;
}