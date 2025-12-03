#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <sstream>

struct range
{
    int64_t from;
    int64_t to;
};

void parse(const std::string &content, std::vector<range> &productIds)
{
    std::stringstream ss(content);
    std::string item;

    while (std::getline(ss, item, ','))
    {
        int32_t start = item.find_first_of('-', 0);
        int64_t from = std::stoll(item.substr(0, start));
        int64_t to = std::stoll(item.substr(start + 1));
        range product_range{from, to};
        productIds.emplace_back(product_range);
    }
}

int32_t number_of_digits(int64_t number)
{
    int32_t count = 0;

    while (number > 0)
    {
        count++;
        number /= 10;
    }

    return count;
}

bool is_valid_prodct(int64_t productId)
{
    int32_t digits = number_of_digits(productId);

    if (digits % 2 != 0)
        return true;

    int64_t key = static_cast<int>(std::pow(10, digits / 2));

    int64_t right = productId % key;
    int64_t left = productId / key;

    return right != left;
}

int64_t sum_of_invalid_ids(std::vector<range> &productIds)
{
    int64_t sum = 0;

    for (auto &productRange : productIds)
    {
        for (int64_t start = productRange.from; start <= productRange.to; start++)
        {
            if (!is_valid_prodct(start))
                sum += start;
        }
    }

    return sum;
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
    std::getline(file, line);
    std::vector<range> productIds;
    parse(line, productIds);

    std::cout << "Sum of Invalid Product Ids : " << sum_of_invalid_ids(productIds) << std::endl;
}