const std = @import("std");
const math = @import("math");
const string = @import("string");

pub fn say(msg: []const u8) void {
    const stdout = std.io.getStdOut().writer();
    _ = stdout.print("{s}\n", .{msg});
}

pub fn goMath(expr: []const u8) i64 {
    return math.eval(expr);
}

pub fn concat(a: []const u8, b: []const u8) []u8 {
    return string.concat(a, b);
}
