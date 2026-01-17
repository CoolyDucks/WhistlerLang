const std = @import("std");

pub fn build(b: *std.Build) void {
    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});

    const exe = b.addExecutable(.{
        .name = "whistlerlang",
        .root_source_file = .{ .path = "source/main.zig" },
        .target = target,
        .optimize = optimize,
    });

    b.installArtifact(exe);
}
