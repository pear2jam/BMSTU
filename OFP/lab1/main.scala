def isPalindrome(lst: List[Int]): Boolean = {
    lst == lst.reverse
}

def sliding(lst: List[Int], size: Int): List[List[Int]] = {
    def sliding_f(lst: List[Int], acc: List[List[Int]]): List[List[Int]] = {
        if (lst.length < size) acc
        else sliding_f(lst.tail, acc :+ lst.take(size))
    }

    sliding_f(lst, List())
}

def subLists(lst: List[Int]): List[List[Int]] = {
    if (lst.isEmpty) List(List())
    else {
        (for {
            i <- 1 to lst.length
            sub <- sliding(lst, i)
        } yield sub).toList
    }
}

def palindromes(lst: List[Int]): List[List[Int]] = {
    subLists(lst).filter(isPalindrome)
}

palindromes(List(1, 2, 1, 2, 2, 1, 3))
