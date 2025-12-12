#include <iostream>
#include <fstream>
#include <sstream>

struct Position
{
    int32_t row;
    int32_t col;
};

Position get_start(const std::vector<std::vector<char>> &grid)
{
    const int32_t rows = grid.size();
    const int32_t cols = grid[0].size();

    for (int32_t i = 0; i < rows; i++)
    {
        for (int32_t j = 0; j < cols; j++)
        {
            if (grid[i][j] == 'S')
            {
                return {i, j};
            }
        }
    }

    return {0, 0};
}

int32_t traverse(std::vector<std::vector<char>> &grid, Position next)
{
    const int32_t rows = grid.size();
    const int32_t cols = grid[0].size();

    if (next.row < 0 || next.row >= rows || next.col >= cols || next.col < 0)
        return 0;

    if (grid[next.row][next.col] == '^')
    {
        return 1 + traverse(grid, {next.row + 1, next.col - 1}) +
               traverse(grid, {next.row + 1, next.col + 1});
    }
    else if (grid[next.row][next.col] == '.')
    {
        grid[next.row][next.col] = '|';
        return traverse(grid, {next.row + 1, next.col});
    }

    return 0;
}

int64_t timelines(std::vector<std::vector<char>> &grid, Position next,
                  std::vector<std::vector<int64_t>> &memo)
{
    const int32_t rows = grid.size();
    const int32_t cols = grid[0].size();

    if (next.row >= rows - 1)
        return memo[next.row][next.col] = 1;

    if (memo[next.row][next.col] != -1)
        return memo[next.row][next.col];

    int64_t ways = 0;

    if (grid[next.row][next.col] == '^')
    {
        ways += timelines(grid, {next.row + 1, next.col - 1}, memo);
        ways += timelines(grid, {next.row + 1, next.col + 1}, memo);
    }
    else
    {
        ways += timelines(grid, {next.row + 1, next.col}, memo);
    }
    memo[next.row][next.col] = ways;
    return ways;
}

void timelines_and_num_beam_splits(std::vector<std::vector<char>> &grid)
{

    int32_t rows = grid.size();
    int32_t cols = grid[0].size();
    Position start = get_start(grid);

    std::vector<std::vector<int64_t>> memo(rows, std::vector<int64_t>(cols, -1));

    // traverse modifies grid, hence timilines calculated first
    std::cout << "(PART 2) Timelines : "
              << timelines(grid, {start.row + 1, start.col}, memo) << std::endl;

    std::cout << "(PART 1) Beam Splits : "
              << traverse(grid, {start.row + 1, start.col}) << std::endl;
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
    std::vector<std::vector<char>> grid;

    while (std::getline(file, line))
    {
        grid.emplace_back(line.begin(), line.end());
    }

    timelines_and_num_beam_splits(grid);
}