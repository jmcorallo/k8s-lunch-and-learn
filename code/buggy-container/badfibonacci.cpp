// Adapted from
// https://stackoverflow.com/questions/16324445/memory-leak-in-code-snippet/16324489

#include <iostream>

class Fibonacci
{
public:
    int m_valuefound;
    int m_tofind;
    long int *m_memo;

    int findValue(int value){
        if (m_memo[value] == 0) {
            if (value == 0 || value == 1) {
                m_memo[value] = 1;
            } else {
                m_memo[value] = findValue(value-1) + findValue(value-2);
            }
        }
        return m_memo[value];
    }

    void setToFind(int value){
        m_tofind = value;
        m_memo = new long int[value];
        std::fill_n(m_memo,value,0);
    }

    void solve(){
        int value = m_tofind;
        int result = findValue(value);
        std::cout<< "Value is: " << result << std::endl;
    }
    ~Fibonacci(){};
};

int main (int argc, char * const argv[]) {
    //std::cout << "Starting memory-leaking fibonacci";
    int count=0;
    while(true){
        count++;
        Fibonacci *fibo = new Fibonacci;
        fibo->setToFind(count);
        fibo->solve();
    }
    return 0;
}