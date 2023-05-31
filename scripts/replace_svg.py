import glob
from xml.etree import ElementTree

def replace_svg():
    html_files = glob.glob("*.html")
    for fl in html_files:
        tree = ElementTree.parse(fl)
        doc = tree.getroot()
        ns = { "xhtml": "http://www.w3.org/1999/xhtml" }
        node = tree.find(".//xhtml:object[@type='image/svg+xml']/..", ns)
        if node is None:
            tree.write(fl)
            continue
        fn = node.find("{http://www.w3.org/1999/xhtml}object[@type='image/svg+xml']").attrib["data"]
        node.remove(node.find("{http://www.w3.org/1999/xhtml}object"))
        svg_tree = ElementTree.parse(fn)
        svg_root = svg_tree.getroot()
        node.insert(0, svg_root)
        tree.write(fl)

if __name__ == "__main__":
    replace_svg()
