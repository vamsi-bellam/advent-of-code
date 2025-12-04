#include <iostream>
#include <fstream>
#include <string>
#include <vector>

int32_t max_joltage(std::string &bank)
{

    int32_t largest_value_index{0};
    size_t bank_size{bank.size()};

    for (size_t i = 0; i < bank_size; i++)
    {
        if (bank[i] > bank[largest_value_index])
        {
            largest_value_index = i;
        }
    }

    int32_t left_index{-1};
    if (largest_value_index > 0)
    {
        int32_t left_start = largest_value_index - 1;
        left_index = left_start;

        while (left_start >= 0)
        {
            if (bank[left_start] > bank[left_index])
            {
                left_index = left_start;
            }
            left_start--;
        }
    }

    int32_t right_index{-1};
    if (largest_value_index < bank_size - 1)
    {
        int32_t right_start = largest_value_index + 1;
        right_index = right_start;

        while (right_start < bank_size)
        {
            if (bank[right_start] > bank[right_index])
            {
                right_index = right_start;
            }
            right_start++;
        }
    }

    int32_t right_value{0}, left_value{0};

    if (left_index != -1)
    {
        left_value = (bank[left_index] - '0') * 10 + (bank[largest_value_index] - '0');
    }

    if (right_index != -1)
    {
        right_value = (bank[largest_value_index] - '0') * 10 + (bank[right_index] - '0');
    }

    return std::max(left_value, right_value);
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
    std::vector<std::string> banks;

    while (std::getline(file, line))
    {
        banks.emplace_back(line);
    }

    int64_t total_joltage{0};

    for (auto &bank : banks)
    {
        total_joltage += max_joltage(bank);
    }

    std::cout << "(PART 1) Total Max Joltage : " << total_joltage << std::endl;
}