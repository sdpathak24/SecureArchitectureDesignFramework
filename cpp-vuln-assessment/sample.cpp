#include <iostream>
#include <cstring>

int main()
{
    char src[] = "This is dangerous!";
    char dest[10];

    // Unsafe use of strcpy - possible buffer overflow
    strcpy(dest, src);

    std::cout << "Copied string: " << dest << std::endl;

    return 0;
}
