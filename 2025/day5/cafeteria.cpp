#include <iostream>
#include <fstream>
#include <sstream>

int main()
{
    std::ifstream file("input.txt");

    if (!file)
    {
        std::cout << "Error opening input file!" << std::endl;
        return 1;
    }

    std::string line;
    std::vector<std::pair<int64_t, int64_t>> fresh_ingredients;
    std::vector<int64_t> queries;

    while (std::getline(file, line))
    {
        if (line.size() == 0)
            break;
        int32_t start = line.find_first_of('-', 0);
        int64_t from = std::stoll(line.substr(0, start));
        int64_t to = std::stoll(line.substr(start + 1));
        fresh_ingredients.emplace_back(from, to);
    }

    while (std::getline(file, line))
    {
        queries.emplace_back(std::stoll(line));
    }

    int64_t total_fresh{0};

    for (auto query : queries)
    {
        for (auto &range : fresh_ingredients)
        {
            if (range.first <= query && query <= range.second)
            {
                total_fresh++;
                break;
            }
        }
    }

    std::cout << "(PART 1) : Total fresh : " << total_fresh << std::endl;
}