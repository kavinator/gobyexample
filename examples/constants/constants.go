// Go поддерживает _константы_ для строк, булевых
// и числовых типов данных

package main

import "fmt"
import "math"

// `const` объявляет переменную как константу.
const s string = "constant"

func main() {
    fmt.Println(s)

    // Оператор `const` может применяться везде,
    // где допустимо применение оператора `var`
    const n = 500000000

    // Выражения с константами выполняют
    // арифметические действия с произвольной точностью.
    const d = 3e20 / n
    fmt.Println(d)

    // Числовые константы не имеют типа,
    // пока он не будет присвоен явным приведением.
    fmt.Println(int64(d))

    // Числу может быть присвоен тип при использовании в контексте,
    // который его требует. Например в случае присвоения
    // переменной или при вызове в функции.
    // Например `math.Sin` ожидает `float64`.
    fmt.Println(math.Sin(n))
}