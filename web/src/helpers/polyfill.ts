(() => {
  if (!String.prototype.replaceAll) {
    String.prototype.replaceAll = function (str: any, newStr: any) {
      if (Object.prototype.toString.call(str).toLocaleLowerCase() === "[Object regexp]") {
        return this.replace(str, newStr);
      }
      return this.replace(new RegExp(str, "g"), newStr);
    };
  }
})();

export default null;
