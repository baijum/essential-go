import glob

images_map = {
        "sphinx/quickstart.rst": {
            "the idea, consider this mathematical function:": [
                "\n.. figure:: _static/function/formula.png\n",
                "\n   Figure 2.1: Mathematical function for area of a circle\n",
                ],
            "the input value the output varies.": [
                "\n.. figure:: _static/function/blackbox.png\n",
                "\n   Figure 2.2: Blackbox representation of a function\n",
                ],
            },
        "sphinx/functions.rst": {
            "the Quick Start chapter.": [
                "\n.. figure:: _static/function/formula.png\n",
                "\n   Figure 5.1: Mathematical function for area of a circle\n",
                ],
            "the input value the output varies.": [
                "\n.. figure:: _static/function/blackbox.png\n",
                "\n   Figure 5.2: Blackbox representation of a function\n",
                ],
            },
        }

def update_code_block():
    rst_files = glob.glob("sphinx/*.rst")
    rst_files.remove("sphinx/index.rst")
    for rst_file in rst_files:
        rst_lines = open(rst_file).readlines()
        new_rst_lines = []
        for line in rst_lines:
            if line.startswith(".. code:: default"):
                new_line = line.replace(".. code:: default", ".. code:: text")
                new_rst_lines.append(new_line)
                continue
            if line.startswith(".. code:: go"):
                new_line = line.replace(".. code:: go", ".. code-block:: go")
                new_rst_lines.append(new_line)
                new_rst_lines.append("   :linenos:\n")
                continue
            found_figure = False
            if rst_file in images_map:
                for k, v in images_map[rst_file].items():
                    if k in line:
                        found_figure = True
                        new_rst_lines.append(line)
                        for l in v:
                            new_rst_lines.append(l)
                        break

            if not found_figure:
                new_rst_lines.append(line)
        final_text = "".join(new_rst_lines)
        with open(rst_file, "w") as fd:
            fd.write(final_text)

if __name__ == "__main__":
    update_code_block()
