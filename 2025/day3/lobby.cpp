#include <iostream>
#include <fstream>
#include <string>
#include <vector>

int32_t max_joltage(const std::string &bank)
{

    int32_t largest_value_index{0};
    size_t bank_size{bank.size()};

    for (int32_t i = 0; i < bank_size; i++)
    {
        if (bank[i] > bank[largest_value_index])
        {
            largest_value_index = i;
        }
    }

    int32_t right_value{0}, left_value{0};

    if (largest_value_index > 0)
    {
        int32_t left_start = largest_value_index - 1;
        int32_t left_index = left_start;

        while (left_start >= 0)
        {
            if (bank[left_start] > bank[left_index])
            {
                left_index = left_start;
            }
            left_start--;
        }
        left_value = (bank[left_index] - '0') * 10 + (bank[largest_value_index] - '0');
    }

    if (largest_value_index < bank_size - 1)
    {
        int32_t right_start = largest_value_index + 1;
        int32_t right_index = right_start;

        while (right_start < bank_size)
        {
            if (bank[right_start] > bank[right_index])
            {
                right_index = right_start;
            }
            right_start++;
        }
        right_value = (bank[largest_value_index] - '0') * 10 + (bank[right_index] - '0');
    }

    return std::max(left_value, right_value);
}

int64_t max_joltage_12(const std::string &bank)
{

    // vector kind of monotonic stack to track largest possible digits
    std::vector<char> bank_values;

    // to ensure vector have at least 12 elements
    int32_t max_removals = bank.size() - 12;

    for (const char c : bank)
    {
        while (!bank_values.empty() && c > bank_values.back() && max_removals > 0)
        {
            bank_values.pop_back();
            max_removals--;
        }

        bank_values.emplace_back(c);
    }

    return std::stoull(std::string(bank_values.begin(), bank_values.begin() + 12));
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

    int64_t total_joltage{0}, total_joltage_12{0};

    for (const auto &bank : banks)
    {
        total_joltage += max_joltage(bank);
        total_joltage_12 += max_joltage_12(bank);
    }

    std::cout << "(PART 1) Total Max Joltage : " << total_joltage << std::endl;
    std::cout << "(PART 2) Total Max Joltage : " << total_joltage_12 << std::endl;
}