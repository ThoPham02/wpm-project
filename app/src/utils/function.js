export const wordsCount = (content) => {
  const str = content
    .replace(/(^\s*)|(\s*$)/gi, "")
    .replace(/[ ]{2,}/gi, " ")
    .replace(/\n /, "\n");
  const words = str.split(" ").length;

  return words;
};

export const lengthCount = (content) => {
  const str = content.trim();
  const length = str.length;
  return length;
};
