const std = @import("std");

pub fn main() !void {
    var it = std.process.args();
    _ = it.next();

    const cmd = it.next() orelse {
        printUsage();
        return;
    };

    const file_name = it.next() orelse {
        printUsage();
        return;
    };

    if (!std.mem.eql(u8, cmd, "run")) {
        printUsage();
        return;
    }

    const file = try std.fs.cwd().openFile(file_name, .{});
    defer file.close();

    const data = try file.readToEndAlloc(std.heap.page_allocator, 1024 * 1024);
    defer std.heap.page_allocator.free(data);

    var lines = std.mem.split(u8, data, "\n");
    while (lines.next()) |line| {
        const trimmed = std.mem.trim(u8, line, " \t\r");
        if (trimmed.len == 0) continue;
        try compileAndRun(trimmed);
    }
}

fn printUsage() void {
    const out = std.io.getStdOut().writer();
    _ = out.print("usage: whistlerlang run <file.whlst>\n", .{}) catch {};
}

fn compileAndRun(line: []const u8) !void {
    if (!std.mem.startsWith(u8, line, "say(")) {
        return error.InvalidSyntax;
    }

    if (!std.mem.endsWith(u8, line, ")")) {
        return error.InvalidSyntax;
    }

    const inner = line[4 .. line.len - 1];
    if (inner.len < 2) return error.InvalidSyntax;
    if (inner[0] != '"' or inner[inner.len - 1] != '"') {
        return error.InvalidSyntax;
    }

    const value = inner[1 .. inner.len - 1];
    try runtimeSay(value);
}

fn runtimeSay(text: []const u8) !void {
    const out = std.io.getStdOut().writer();
    try out.print("{s}\n", .{text});
}
