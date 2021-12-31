def _impl(ctx):
    args = [f.path for f in ctx.files.srcs]
    flags = [
        "archive",
        "--output=%s" % ctx.outputs.output.path,
        "--prefix=%s" % ctx.attr.package_dir,
    ]

    ctx.actions.run(
        inputs = ctx.files.srcs,
        outputs = [ctx.outputs.output],
        arguments = flags + args,
        progress_message = "Archiving into %s" % ctx.outputs.output.short_path,
        executable = ctx.executable._tar_tool,
    )

_pkg_tar = rule(
    implementation = _impl,
    attrs = {
        "srcs": attr.label_list(allow_files = True),
        "package_dir": attr.string(default = "/"),
        "output": attr.output(mandatory = True),
        "_tar_tool": attr.label(
            executable = True,
            cfg = "exec",
            allow_files = True,
            default = Label("//cmd:tarchiver"),
        ),
    },
)

def pkg_tar(output = "", **kwargs):
    _pkg_tar(
        output = "%s.tar" % kwargs["name"],
        **kwargs
    )

def _file_size_impl(ctx):
    runfiles_path = "$0.runfiles/"
    data_file_root = runfiles_path + ctx.workspace_name + "/"

    _runfiles = ctx.runfiles(
        files = [
            ctx.files._size_tool[0],
            ctx.files.file[0],
        ],
        root_symlinks = {
            "tool_dep": ctx.files._size_tool[0],
            "data_dep": ctx.files.file[0],
        },
    )

    ctx.actions.write(
        output = ctx.outputs.executable,
        content = "#!/usr/bin/env bash\n%s size %s" % (
            runfiles_path + "tool_dep",
            runfiles_path + "data_dep",
        ),
        is_executable = True,
    )

    return [DefaultInfo(
        runfiles = _runfiles,
    )]

file_size = rule(
    implementation = _file_size_impl,
    attrs = {
        "file": attr.label(
            mandatory = True,
        ),
        "_size_tool": attr.label(
            executable = True,
            cfg = "exec",
            allow_files = True,
            default = Label("//cmd:tarchiver"),
        ),
    },
    executable = True,
)
