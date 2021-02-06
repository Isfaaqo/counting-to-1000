#include <iostream>
using namespace std;
int main() {
	int c = 0;
	cout << c;
	while(c < 1000) {
		if (cin.get() == '\n') {
			c++;
			cout << c;
		}
	}
}