#include <iostream>
#include <fstream>
#include <sstream>
#include <numeric>
#include <functional>

int main()
{
    std::ifstream file("input.txt");

    if (!file)
    {
        std::cout << "Error opening input file!" << std::endl;
        return 1;
    }

    std::string line;
    std::vector<std::string> input;
    std::vector<std::vector<int64_t>> numbers;

    while (std::getline(file, line))
    {
        input.emplace_back(line);
    }

    for (size_t i = 0; i < input.size() - 1; i++)
    {
        std::string num;
        std::istringstream iss(input.at(i));
        size_t j{0};

        while (iss >> num)
        {
            if (numbers.size() < j + 1)
            {
                numbers.emplace_back(std::vector<int64_t>{std::stoll(num)});
            }
            else
            {
                numbers.at(j).emplace_back(std::stoll(num));
            }

            j++;
        }
    }

    int64_t total_sum{0}, index{0};
    std::string op;
    std::vector<std::string> ops;
    std::istringstream iss(input.back());

    while (iss >> op)
    {
        ops.emplace_back(op);
        std::vector<int64_t> &nums = numbers.at(index);
        if (op == "+")
        {
            total_sum += std::accumulate(nums.begin(), nums.end(), int64_t{0});
        }
        else if (op == "*")
        {
            total_sum += std::accumulate(nums.begin(), nums.end(), int64_t{1},
                                         std::multiplies<int64_t>());
        }
        index++;
    }

    std::cout << "(PART 1) : Total sum : " << total_sum << std::endl;
}