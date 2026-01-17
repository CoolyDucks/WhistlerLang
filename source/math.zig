const std = @import("std");

pub fn eval(expr: []const u8) i64 {
    var result: i64 = 0;
    var op: u8 = 0;
    var num: i64 = 0;
    for (expr) |c| {
        if (c >= '0' and c <= '9') {
            num = num * 10 + (@intCast(i64, c - '0'));
        } else if (c == '+' or c == '-' or c == '*' or c == '/') {
            if (op == 0) result = num;
            else if (op == '+') result += num;
            else if (op == '-') result -= num;
            else if (op == '*') result *= num;
            else if (op == '/') result /= num;
            op = c;
            num = 0;
        }
    }
    if (op == 0) result = num;
    else if (op == '+') result += num;
    else if (op == '-') result -= num;
    else if (op == '*') result *= num;
    else if (op == '/') result /= num;
    return result;
}
