#include <iostream>
#include <fstream>
#include <string>

int main()
{
    std::ifstream file("input.txt");

    if (!file)
    {
        std::cout << "Error opening input file!" << std::endl;
        return 1;
    }

    std::string line;
    int32_t passowrd{0}, value{0}, new_password{0};
    std::string dir;
    int64_t current{50};

    while (std::getline(file, line))
    {
        std::cout << "line : " << line << std::endl;

        dir = line.substr(0, 1);
        auto value = std::stoll(line.substr(1));

        new_password += value / 100;
        value %= 100;

        std::cout << "value : " << value << std::endl;

        if (dir == "L")
        {
            // we can cross 0, if we are going more than current value on left.
            /*
              current should be > 0 because, if it is 0 means we already
              counted in before iteration
            */
            if (current > 0 && value >= current)
            {
                new_password++;
            }

            current = current - value;
            current = current < 0 ? current + 100 : current;
        }
        else if (dir == "R")
        {

            current = current + value;
            // if current crosses 100 then we crossed 0
            if (current >= 100)
            {
                new_password++;
            }
            current = current > 99 ? current - 100 : current;
        }
        else
        {
            std::cout << "Malformed input!" << std::endl;
            return 1;
        }

        if (current == 0)
            passowrd++;
    }

    std::cout << "(PART1) PASSWORD : " << passowrd << std::endl;

    std::cout << "(PART2) PASSWORD : " << new_password << std::endl;
}