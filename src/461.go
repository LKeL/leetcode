func hammingDistance(x int, y int) int {
    result := x ^ y;
    a := 0;
    for {
        if result & 1 == 1 {
            a++;
        }
        result >>= 1;
        if result == 0 {
            break;
        }
    }
    return a;
}
