const std = @import("std");

pub fn concat(a: []const u8, b: []const u8) []u8 {
    var result = std.heap.page_allocator.alloc(u8, a.len + b.len) catch return &[_]u8{};
    for (a) |c, i| result[i] = c;
    for (b) |c, i| result[a.len + i] = c;
    return result;
}
