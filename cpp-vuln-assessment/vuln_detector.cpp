#include <iostream>
#include <fstream>
#include <regex>
#include <string>

int main()
{
    std::ifstream infile("sample.cpp");
    std::string line;
    int line_no = 0, issues = 0;
    std::regex strcpy_pattern("strcpy\\s*\\(");

    while (getline(infile, line))
    {
        ++line_no;
        if (std::regex_search(line, strcpy_pattern))
        {
            std::cout << "Vulnerability at line " << line_no << ": strcpy usage\n";
            ++issues;
        }
    }
    std::cout << "Scan complete: " << issues << " issues found\n";
    return 0;
}
