function filter(
  list = [],
  conditions = {
    type: "string",
    condition: "" | /()/ | function (k) {} | [],
    positive: false,
  }
) {
  const { type, condition, positive } = conditions;
  switch (type) {
    case "string":
      if (typeof condition != "string") break;
      return list.filter((v) =>
        positive ? v.match(condition) : !v.match(condition)
      );
    case "function":
      if (typeof condition != "function") break;
      return list.filter(condition);
    default:
      return list;
  }
  return list;
}
