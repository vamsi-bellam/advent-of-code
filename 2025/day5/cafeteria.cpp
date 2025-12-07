#include <iostream>
#include <fstream>
#include <sstream>
#include <algorithm>

using IngredientsRange = std::pair<int64_t, int64_t>;
using FreshIngredients = std::vector<IngredientsRange>;

int64_t get_total_fresh(FreshIngredients &fresh_ingredients)
{
    if (fresh_ingredients.empty())
        return 0;

    IngredientsRange current{fresh_ingredients.at(0)};

    int64_t total_fresh{0};

    for (const auto &next : fresh_ingredients)
    {
        if (next.first <= current.second)
        {
            current.second = std::max(current.second, next.second);
        }
        else
        {
            total_fresh += current.second - current.first + 1;
            current = next;
        }
    }

    total_fresh += current.second - current.first + 1;

    return total_fresh;
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
    FreshIngredients fresh_ingredients;
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

    std::sort(fresh_ingredients.begin(),
              fresh_ingredients.end(),
              [](IngredientsRange &a, IngredientsRange &b)
              { return a.first < b.first; });

    std::cout << "(PART 2) : Total fresh : "
              << get_total_fresh(fresh_ingredients) << std::endl;
}