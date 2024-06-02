"use strict";
const common_vendor = require("../../common/vendor.js");
var vk;
var goodsCache = {};
const _sfc_main = {
  name: "vk-data-goods-sku-popup",
  emits: [
    "update:modelValue",
    "input",
    "update-goods",
    "open",
    "close",
    "add-cart",
    "buy-now",
    "cart",
    "buy",
    "num-change"
  ],
  props: {
    // true 组件显示 false 组件隐藏
    value: {
      Type: Boolean,
      default: false
    },
    modelValue: {
      Type: Boolean,
      default: false
    },
    // vk云函数路由模式参数开始-----------------------------------------------------------
    // 商品id
    goodsId: {
      Type: String,
      default: ""
    },
    // vk路由模式框架下的云函数地址
    action: {
      Type: String,
      default: ""
    },
    // vk云函数路由模式参数结束-----------------------------------------------------------
    // 该商品已抢完时的按钮文字
    noStockText: {
      Type: String,
      default: "该商品已抢完"
    },
    // 库存文字
    stockText: {
      Type: String,
      default: "库存"
    },
    // 商品表id的字段名
    goodsIdName: {
      Type: String,
      default: "_id"
    },
    // sku表id的字段名
    skuIdName: {
      Type: String,
      default: "_id"
    },
    // sku_list的字段名
    skuListName: {
      Type: String,
      default: "sku_list"
    },
    // spec_list的字段名
    specListName: {
      Type: String,
      default: "spec_list"
    },
    // 库存的字段名 默认 stock
    stockName: {
      Type: String,
      default: "stock"
    },
    // sku组合路径的字段名
    skuArrName: {
      Type: String,
      default: "sku_name_arr"
    },
    // 默认单规格时的规格组名称
    defaultSingleSkuName: {
      Type: String,
      default: "默认"
    },
    // 模式 1:都显示  2:只显示购物车 3:只显示立即购买 4:显示缺货按钮 默认 1
    mode: {
      Type: Number,
      default: 1
    },
    // 点击遮罩是否关闭组件 true 关闭 false 不关闭 默认true
    maskCloseAble: {
      Type: Boolean,
      default: true
    },
    // 顶部圆角值
    borderRadius: {
      Type: [String, Number],
      default: 0
    },
    // 商品缩略图字段名(未选择sku时)
    goodsThumbName: {
      Type: [String],
      default: "goods_thumb"
    },
    // 商品缩略图背景颜色，如#999999
    goodsThumbBackgroundColor: {
      Type: String,
      default: "transparent"
    },
    // 最小购买数量 默认 1
    minBuyNum: {
      Type: [Number, String],
      default: 1
    },
    // 最大购买数量 默认 100000
    maxBuyNum: {
      Type: [Number, String],
      default: 1e5
    },
    // 步进器步长 默认 1
    stepBuyNum: {
      Type: [Number, String],
      default: 1
    },
    // 是否只能输入 step 的倍数
    stepStrictly: {
      Type: Boolean,
      default: false
    },
    // 自定义获取商品信息的函数,支付宝小程序不支持该属性,请使用localdata属性
    customAction: {
      Type: [Function],
      default: null
    },
    // 本地数据源
    localdata: {
      type: Object
    },
    // 价格的字体颜色
    priceColor: {
      Type: String
    },
    // 立即购买按钮的文字
    buyNowText: {
      Type: String,
      default: "立即购买"
    },
    // 立即购买按钮的字体颜色
    buyNowColor: {
      Type: String
    },
    // 立即购买按钮的背景颜色
    buyNowBackgroundColor: {
      Type: String
    },
    // 加入购物车按钮的文字
    addCartText: {
      Type: String,
      default: "加入购物车"
    },
    // 加入购物车按钮的字体颜色
    addCartColor: {
      Type: String
    },
    // 加入购物车按钮的背景颜色
    addCartBackgroundColor: {
      Type: String
    },
    // 不可点击时,按钮的样式
    disableStyle: {
      Type: Object,
      default: null
    },
    // 按钮点击时的样式
    activedStyle: {
      Type: Object,
      default: null
    },
    // 按钮常态的样式
    btnStyle: {
      Type: Object,
      default: null
    },
    // 是否显示右上角关闭按钮
    showClose: {
      Type: Boolean,
      default: true
    },
    // 关闭按钮的图片地址 https://img.alicdn.com/imgextra/i1/121022687/O1CN01ImN0O11VigqwzpLiK_!!121022687.png
    closeImage: {
      Type: String,
      default: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACoAAAAqCAYAAADFw8lbAAAEyUlEQVR42sSZeWwNURTGp4OqtBo7sSXELragdkpQsRRJ1Zr4hyJiJ9YgxNIg1qANiT+E1i5IY0kVVWtQEbuEKLFGUSH27/ANN5PXmTvzupzkl/tm8t6b7517lnvvC0lKSjJ8WmnQAUSDFqABqALKgl8gD7wE90E2SAeXwFf1SxISErQeVtKHwCgwFsSDSIf3hYFKoCkYDBaDdyAViHdueHmoF6FtwDLQ23b/E7gM7oIcejIERIDaoBFoC8qA8mA8SQNz6W1XC9GY+nCQCCYAk/c+gF0gBZwH312+IxR0BCPBUIaH2A+wHsxHCHxx+gLT5QGN6a2JfG8uvVCDws9oiDQYlxkMGfHyQvARlADTwcXk5OT6foV2kS8ATXidymlcyen1a/Jjl9IJh3hPkjELYqO8Cu0KjjNZvtETw5jFBWXPmGSTGQKSeOn5iQ0kVLL0CINfPNcPbDMKyRCbGzEMBJ+ZD8cChYFdqGTqfsWT8otPGoVsEHsMwxDFs3shNsxJ6BrQ0Po8OGUUkVHsNCVml+cntB1jUWwn2GEUsTEMrASbDK+2CCQ0kYX6nfLLisMmKqUr0S60M+jG10vAm+JSCa8+x7CKlzHwaktV6DiObzUzPJIxFO1BQ12wGtTReO9GetVgY/kjNJzZbcWmTjHfxw51AsRqvL8eOAtmsJuFu3g1l+1ZLB5eDTVZ3K0P7tL0TkWOpSg61kVkBtuuNRthGs+wtJST5aQI7cEbkkRXNYVKgX6kIdYuUhYzMQwxN8tiExCLFqHNeSF9/aem0BzGp5PYQCJ7c/Gsk1RfuSD6U1dNpcDf9ZigTmKbMRZ9iVTsHscGJluW2FMf1SSQWGnBmaB6kCJVTVVNJZE++Cx9drEllS1KMCINpURFmEbBWA63Fz9s95cGIdJgp/zXmT4pZcOvSUzuZttTbblmnc3PIjjmidDXvKgdhMh0JdbzuCjWrbNOVovjS5P7bkPJ/mBESkz2BO0166ybNeJ431S2q+01NntuIq3E0amzjiZtk9tssWyTDzO4525bACK9NAUn68TtkNhpEXpOSagRml+S6iLSSeweHv242Qhl13rRyvoDvDlKyTQny/ZQJ+1iH7vVbEx7OR5UiKVIO7VicgvHCtwrudloMIV7/0uadVYW57O4Wvvi8v4pymlKkrpwvsDeLLZAY2pkwbAB3PSQfC+4cH7l4k1ZH8zkZRq8ecO+Z5rN40JJqnXFuGfaxPCTLjcn0OZOpnArXw8HY4paIbw5CcMgXq6HN2/mt6+XGLrN15tBryIUGavMpCTrfKcDCKkAceA9S8nhAOehhSUyhXpkBxxnP4YM1InugP7cBkjBPcqVUWFYCEROxXiQz5JlXV+IfKh7mpfJac+lZ6V87QXVClBkTc7YWsWTPSDyitfzUTlJlj8TbvE6jluDOdwZ+jX57GLO3ADeuyZrDYi86vV81FD2UVGsmT+5Zl0BnkhoseOEaogL46pqO4v/IqUEyalIR4h85BgjHv6+aUWRMbb7EstX6O0cpT1Gco0ry8fWygLDMjmDnQeBt3Qe7uVfkeugDwVLcsVzGsuwLXbV+I63XNAkG5r/hvgRqgqWs6pJPKrsbvz/Q6yyun0w/h6lP+BnzrCpfPMT2L8FGAA7k1GZ/vnaqAAAAABJRU5ErkJggg=="
    },
    // 是否隐藏库存显示
    hideStock: {
      Type: Boolean,
      default: false
    },
    // 颜色主题
    theme: {
      Type: String,
      default: "default"
    },
    // 请求中的提示
    actionTips: {
      Type: String,
      default: "请求中..."
    },
    // 默认选中的SKU
    defaultSelect: {
      Type: Object
    },
    // 是否使用缓存
    useCache: {
      Type: Boolean,
      default: true
    },
    /**
     * 默认商品,设置该值可快速展示商品
     * 逻辑: 先展示 defaultGoods 信息,再取数据库,再更新页面(通常为更新库存)
     */
    defaultGoods: {
      Type: Object
    },
    /**
     * 金额是否需要除以100
     * 1:金额会除以100
     * 0:金额不会除以100
     */
    amountType: {
      Type: Number,
      default: 1
    },
    // 每次选择完SKU后，购买数量归1，如果有最小购买数量，则设置为最小购买数量
    selectedInit: {
      Type: Boolean,
      default: false
    },
    // 是否开启底部安全区适配，默认true
    safeAreaInsetBottom: {
      Type: Boolean,
      default: true
    }
  },
  data() {
    return {
      safeBottom: 0,
      // 留出底部安全距离
      complete: false,
      // 组件是否加载完成
      goodsInfo: {},
      // 商品信息
      isShow: false,
      // true 显示 false 隐藏
      initKey: true,
      // 是否需要初始化 true 是 false 否
      shopItemInfo: {},
      // 存放要和选中的值进行匹配的数据（因百度小程序setData不支持中文字段，故不编译shopItemInfo变量）
      selectArr: [],
      // 存放被选中的值
      subIndex: [],
      // 是否选中 因为不确定是多规格还是单规格，所以这里定义数组来判断
      selectShop: {},
      // 存放最后选中的商品
      selectNum: this.minBuyNum || 1,
      // 选中数量
      outFoStock: false,
      // 是否全部sku都缺货
      openTime: 0,
      themeColor: {
        // 默认主题
        default: {
          priceColor: "rgb(254, 86, 10)",
          buyNowColor: "#ffffff",
          buyNowBackgroundColor: "rgb(254, 86, 10)",
          addCartColor: "#ffffff",
          addCartBackgroundColor: "rgb(255, 148, 2)",
          btnStyle: {
            color: "#333333",
            borderColor: "#f4f4f4",
            backgroundColor: "#ffffff"
          },
          activedStyle: {
            color: "rgb(254, 86, 10)",
            borderColor: "rgb(254, 86, 10)",
            backgroundColor: "rgba(254,86,10,0.1)"
          },
          disableStyle: {
            color: "#c3c3c3",
            borderColor: "#f6f6f6",
            backgroundColor: "#f6f6f6"
          }
        },
        // 红黑主题
        "red-black": {
          priceColor: "rgb(255, 68, 68)",
          buyNowColor: "#ffffff",
          buyNowBackgroundColor: "rgb(255, 68, 68)",
          addCartColor: "#ffffff",
          addCartBackgroundColor: "rgb(85, 85, 85)",
          activedStyle: {
            color: "rgb(255, 68, 68)",
            borderColor: "rgb(255, 68, 68)",
            backgroundColor: "rgba(255,68,68,0.1)"
          }
        },
        // 黑白主题
        "black-white": {
          priceColor: "rgb(47, 47, 52)",
          buyNowColor: "#ffffff",
          buyNowBackgroundColor: "rgb(47, 47, 52)",
          addCartColor: "rgb(47, 47, 52)",
          addCartBackgroundColor: "rgb(235, 236, 242)",
          // btnStyle:{
          // 	color:"rgb(47, 47, 52)",
          // 	borderColor:"rgba(235,236,242,0.5)",
          // 	backgroundColor:"rgba(235,236,242,0.5)",
          // },
          activedStyle: {
            color: "rgb(47, 47, 52)",
            borderColor: "rgba(47,47,52,0.12)",
            backgroundColor: "rgba(47,47,52,0.12)"
          }
        },
        // 咖啡色主题
        coffee: {
          priceColor: "rgb(195, 167, 105)",
          buyNowColor: "#ffffff",
          buyNowBackgroundColor: "rgb(195, 167, 105)",
          addCartColor: "rgb(195, 167, 105)",
          addCartBackgroundColor: "rgb(243, 238, 225)",
          activedStyle: {
            color: "rgb(195, 167, 105)",
            borderColor: "rgb(195, 167, 105)",
            backgroundColor: "rgba(195, 167, 105,0.1)"
          }
        },
        // 浅绿色主题
        green: {
          priceColor: "rgb(99, 190, 114)",
          buyNowColor: "#ffffff",
          buyNowBackgroundColor: "rgb(99, 190, 114)",
          addCartColor: "rgb(99, 190, 114)",
          addCartBackgroundColor: "rgb(225, 244, 227)",
          activedStyle: {
            color: "rgb(99, 190, 114)",
            borderColor: "rgb(99, 190, 114)",
            backgroundColor: "rgba(99, 190, 114,0.1)"
          }
        }
      }
    };
  },
  created() {
    let that = this;
    vk = that.vk;
    if (that.valueCom) {
      that.open();
    }
    const { safeAreaInsets } = common_vendor.index.getSystemInfoSync();
    that.safeBottom = safeAreaInsets.bottom;
  },
  mounted() {
  },
  methods: {
    // 初始化
    init(notAutoClick) {
      let that = this;
      that.selectArr = [];
      that.subIndex = [];
      that.selectShop = {};
      that.selectNum = that.minBuyNum || 1;
      that.outFoStock = false;
      that.shopItemInfo = {};
      let specListName = that.specListName;
      that.goodsInfo[specListName].map((item) => {
        that.selectArr.push("");
        that.subIndex.push(-1);
      });
      that.checkItem();
      that.checkInpath(-1);
      if (!notAutoClick)
        that.autoClickSku();
    },
    // 使用vk路由模式框架获取商品信息
    findGoodsInfo(obj = {}) {
      let that = this;
      let { useCache } = obj;
      if (typeof vk == "undefined") {
        that.toast("custom-action必须是function", "none");
        return false;
      }
      let { actionTips } = that;
      let actionTitle = "";
      let actionAoading = false;
      if (actionTips !== "custom") {
        actionTitle = useCache ? "" : "请求中...";
      } else {
        actionAoading = useCache ? false : true;
      }
      vk.callFunction({
        url: that.action,
        title: actionTitle,
        loading: actionAoading,
        data: {
          goods_id: that.goodsId
        },
        success(data) {
          that.updateGoodsInfo(data.goodsInfo);
          goodsCache[that.goodsId] = data.goodsInfo;
          that.$emit("update-goods", data.goodsInfo);
        },
        fail() {
          that.updateValue(false);
        }
      });
    },
    updateValue(value) {
      let that = this;
      if (value) {
        that.$emit("open", true);
        that.$emit("input", true);
        that.$emit("update:modelValue", true);
      } else {
        that.$emit("input", false);
        that.$emit("close", "close");
        that.$emit("update:modelValue", false);
      }
    },
    // 更新商品信息(库存、名称、图片)
    updateGoodsInfo(goodsInfo) {
      let that = this;
      let { skuListName } = that;
      if (JSON.stringify(that.goodsInfo) === "{}" || that.goodsInfo[that.goodsIdName] !== goodsInfo[that.goodsIdName]) {
        that.goodsInfo = goodsInfo;
        that.initKey = true;
      } else {
        that.goodsInfo[skuListName] = goodsInfo[skuListName];
      }
      if (that.initKey) {
        that.initKey = false;
        that.init();
      }
      let select_sku_info = that.getListItem(
        that.goodsInfo[skuListName],
        that.skuIdName,
        that.selectShop[that.skuIdName]
      );
      Object.assign(that.selectShop, select_sku_info);
      that.defaultSelectSku();
      that.complete = true;
    },
    async open() {
      let that = this;
      that.openTime = new Date().getTime();
      let findGoodsInfoRun = true;
      that.skuListName;
      let useCache = false;
      let goodsInfo = goodsCache[that.goodsId];
      if (goodsInfo && that.useCache) {
        useCache = true;
        that.updateGoodsInfo(goodsInfo);
      } else {
        that.complete = false;
      }
      if (that.customAction && typeof that.customAction === "function") {
        try {
          goodsInfo = await that.customAction({
            useCache,
            goodsId: that.goodsId,
            goodsInfo,
            close: function() {
              setTimeout(function() {
                that.close();
              }, 500);
            }
          }).catch((err) => {
            setTimeout(function() {
              that.close();
            }, 500);
          });
        } catch (err) {
          let { message = "" } = err;
          if (message.indexOf(".catch is not a function") > -1) {
            that.toast("custom-action必须返回一个Promise", "none");
            setTimeout(function() {
              that.close();
            }, 500);
            return false;
          }
        }
        goodsCache[that.goodsId] = goodsInfo;
        if (goodsInfo && typeof goodsInfo == "object" && JSON.stringify(goodsInfo) != "{}") {
          findGoodsInfoRun = false;
          that.updateGoodsInfo(goodsInfo);
          that.updateValue(true);
        } else {
          that.toast("未获取到商品信息", "none");
          that.$emit("input", false);
          return false;
        }
      } else if (typeof that.localdata !== "undefined" && that.localdata !== null) {
        goodsInfo = that.localdata;
        if (goodsInfo && typeof goodsInfo == "object" && JSON.stringify(goodsInfo) != "{}") {
          findGoodsInfoRun = false;
          that.updateGoodsInfo(goodsInfo);
          that.updateValue(true);
        } else {
          that.toast("未获取到商品信息", "none");
          that.$emit("input", false);
          return false;
        }
      } else {
        if (findGoodsInfoRun)
          that.findGoodsInfo({ useCache });
      }
    },
    // 监听 - 弹出层收起
    close(s) {
      let that = this;
      if (new Date().getTime() - that.openTime < 400) {
        return false;
      }
      if (s == "mask") {
        if (that.maskCloseAble !== false) {
          that.$emit("input", false);
          that.$emit("close", "mask");
          that.$emit("update:modelValue", false);
        }
      } else {
        that.$emit("input", false);
        that.$emit("close", "close");
        that.$emit("update:modelValue", false);
      }
    },
    moveHandle() {
    },
    // sku按钮的点击事件
    skuClick(value, index1, index2) {
      let that = this;
      if (value.ishow) {
        if (that.selectArr[index1] != value.name) {
          that.$set(that.selectArr, index1, value.name);
          that.$set(that.subIndex, index1, index2);
        } else {
          that.$set(that.selectArr, index1, "");
          that.$set(that.subIndex, index1, -1);
        }
        that.checkInpath(index1);
        that.checkSelectShop();
      }
    },
    // 检测是否已经选完sku
    checkSelectShop() {
      let that = this;
      if (that.selectArr.every((item) => item != "")) {
        that.selectShop = that.shopItemInfo[that.getArrayToSting(that.selectArr)];
        let stock = that.selectShop[that.stockName];
        if (typeof stock !== "undefined" && that.selectNum > stock) {
          that.selectNum = stock;
        }
        if (that.selectNum > that.maxBuyNum) {
          that.selectNum = that.maxBuyNum;
        }
        if (that.selectNum < that.minBuyNum) {
          that.selectNum = that.minBuyNum;
        }
        if (that.selectedInit) {
          that.selectNum = that.minBuyNum || 1;
        }
      } else {
        that.selectShop = {};
      }
    },
    // 检查路径
    checkInpath(clickIndex) {
      let that = this;
      let specListName = that.specListName;
      let specList = that.goodsInfo[specListName];
      for (let i = 0, len = specList.length; i < len; i++) {
        if (i == clickIndex) {
          continue;
        }
        let len2 = specList[i].list.length;
        for (let j = 0; j < len2; j++) {
          if (that.subIndex[i] != -1 && j == that.subIndex[i]) {
            continue;
          }
          let choosed_copy = [...that.selectArr];
          that.$set(choosed_copy, i, specList[i].list[j].name);
          let choosed_copy2 = choosed_copy.filter(
            (item) => item !== "" && typeof item !== "undefined"
          );
          if (that.shopItemInfo.hasOwnProperty(that.getArrayToSting(choosed_copy2))) {
            specList[i].list[j].ishow = true;
          } else {
            specList[i].list[j].ishow = false;
          }
        }
      }
      that.$set(that.goodsInfo, specListName, specList);
    },
    // 计算sku里面规格形成路径
    checkItem() {
      let that = this;
      let { stockName } = that;
      let skuListName = that.skuListName;
      let originalSkuList = that.goodsInfo[skuListName];
      let skuList = [];
      let stockNum = 0;
      originalSkuList.map((skuItem, index) => {
        if (skuItem[stockName] > 0) {
          skuList.push(skuItem);
          stockNum += skuItem[stockName];
        }
      });
      if (stockNum <= 0) {
        that.outFoStock = true;
      }
      skuList.reduce(
        (arrs, items) => {
          return arrs.concat(
            items[that.skuArrName].reduce(
              (arr, item) => {
                return arr.concat(
                  arr.map((item2) => {
                    if (!that.shopItemInfo.hasOwnProperty(that.getArrayToSting([...item2, item]))) {
                      that.shopItemInfo[that.getArrayToSting([...item2, item])] = items;
                    }
                    return [...item2, item];
                  })
                );
              },
              [[]]
            )
          );
        },
        [[]]
      );
    },
    getArrayToSting(arr) {
      let str = "";
      arr.map((item, index) => {
        item = item.replace(/\./g, "。");
        if (index == 0) {
          str += item;
        } else {
          str += "," + item;
        }
      });
      return str;
    },
    // 检测sku选项是否已全部选完,且有库存
    checkSelectComplete(obj = {}) {
      let that = this;
      let clickTime = new Date().getTime();
      if (that.clickTime && clickTime - that.clickTime < 400) {
        return false;
      }
      that.clickTime = clickTime;
      let { selectShop, selectNum, stockText, stockName } = that;
      if (!selectShop || !selectShop[that.skuIdName]) {
        that.toast("请先选择对应规格", "none");
        return false;
      }
      if (selectNum <= 0) {
        that.toast("购买数量必须>0", "none");
        return false;
      }
      if (selectNum > selectShop[stockName]) {
        that.toast(stockText + "不足", "none");
        return false;
      }
      if (typeof obj.success == "function")
        obj.success(selectShop);
    },
    // 加入购物车
    addCart() {
      let that = this;
      that.checkSelectComplete({
        success: function(selectShop) {
          selectShop.buy_num = that.selectNum;
          that.$emit("add-cart", selectShop);
          that.$emit("cart", selectShop);
        }
      });
    },
    // 立即购买
    buyNow() {
      let that = this;
      that.checkSelectComplete({
        success: function(selectShop) {
          selectShop.buy_num = that.selectNum;
          that.$emit("buy-now", selectShop);
          that.$emit("buy", selectShop);
        }
      });
    },
    // 弹窗
    toast(title, icon) {
      common_vendor.index.showToast({
        title,
        icon
      });
    },
    // 获取对象数组中的某一个item,根据指定的键值
    getListItem(list, key, value) {
      let item;
      for (let i in list) {
        if (typeof value == "object") {
          if (JSON.stringify(list[i][key]) === JSON.stringify(value)) {
            item = list[i];
            break;
          }
        } else {
          if (list[i][key] === value) {
            item = list[i];
            break;
          }
        }
      }
      return item;
    },
    getListIndex(list, key, value) {
      let index = -1;
      for (let i = 0; i < list.length; i++) {
        if (list[i][key] === value) {
          index = i;
          break;
        }
      }
      return index;
    },
    // 自动选择sku前提是只有一组sku,默认自动选择最前面的有库存的sku
    autoClickSku() {
      let that = this;
      let { stockName } = that;
      let skuList = that.goodsInfo[that.skuListName];
      let specListArr = that.goodsInfo[that.specListName];
      if (specListArr.length == 1) {
        let specList = specListArr[0].list;
        for (let i = 0; i < specList.length; i++) {
          let sku = that.getListItem(skuList, that.skuArrName, [specList[i].name]);
          if (sku && sku[stockName] > 0) {
            that.skuClick(specList[i], 0, i);
            break;
          }
        }
      }
    },
    // 主题颜色
    themeColorFn(name) {
      let that = this;
      let { theme, themeColor } = that;
      let color = that[name] ? that[name] : themeColor[theme][name];
      return color;
    },
    defaultSelectSku() {
      let that = this;
      let { defaultSelect } = that;
      if (defaultSelect && defaultSelect.sku && defaultSelect.sku.length > 0) {
        that.selectSku(defaultSelect);
      }
    },
    /**
       * 主动方法 - 设置sku
      that.$refs.skuPopup.selectSku({
        sku:["红色","256G","公开版"],
        num:5
      });
       */
    selectSku(obj = {}) {
      let that = this;
      let { sku: skuArr, num: selectNum } = obj;
      let specListArr = that.goodsInfo[that.specListName];
      if (skuArr && specListArr.length === skuArr.length) {
        let skuClickArr = [];
        let clickKey = true;
        for (let index = 0; index < skuArr.length; index++) {
          let skuName = skuArr[index];
          let specList = specListArr[index].list;
          let index1 = index;
          let index2 = that.getListIndex(specList, "name", skuName);
          if (index2 == -1) {
            clickKey = false;
            break;
          }
          skuClickArr.push({
            spec: specList[index2],
            index1,
            index2
          });
        }
        if (clickKey) {
          that.init(true);
          skuClickArr.map((item) => {
            that.skuClick(item.spec, item.index1, item.index2);
          });
        }
      }
      if (selectNum > 0)
        that.selectNum = selectNum;
    },
    priceFilter(n = 0) {
      let that = this;
      if (typeof n == "string") {
        n = parseFloat(n);
      }
      if (that.amountType === 0) {
        return n.toFixed(2);
      } else {
        return (n / 100).toFixed(2);
      }
    },
    pushGoodsCache(goodsInfo) {
      let that = this;
      let { goodsIdName } = that;
      goodsCache[goodsInfo[goodsIdName]] = goodsInfo;
    },
    // 用于阻止冒泡
    stop() {
    },
    // 图片预览
    previewImage() {
      let that = this;
      let { selectShop, goodsInfo, goodsThumbName } = that;
      let src = selectShop.image ? selectShop.image : goodsInfo[goodsThumbName];
      if (src) {
        common_vendor.index.previewImage({
          urls: [src]
        });
      }
    },
    getMaxStock() {
      let maxStock = 0;
      let that = this;
      let { selectShop = {}, goodsInfo = {}, skuListName, stockName } = that;
      if (selectShop[stockName]) {
        maxStock = selectShop[stockName];
      } else {
        let skuList = goodsInfo[skuListName];
        if (skuList && skuList.length > 0) {
          let valueArr = [];
          skuList.map((skuItem, index) => {
            valueArr.push(skuItem[stockName]);
          });
          let max = Math.max(...valueArr);
          maxStock = max;
        }
      }
      return maxStock;
    },
    numChange(e) {
      this.$emit("num-change", e.value);
    }
  },
  // 计算属性
  computed: {
    valueCom() {
      return this.modelValue;
    },
    // 最大购买数量
    maxBuyNumCom() {
      let that = this;
      let maxStock = that.getMaxStock();
      let max = that.maxBuyNum || 1e5;
      if (max > maxStock) {
        max = maxStock;
      }
      return max;
    },
    // 是否是多规格
    isManyCom() {
      let that = this;
      let { goodsInfo, defaultSingleSkuName, specListName } = that;
      let isMany = true;
      if (goodsInfo[specListName] && goodsInfo[specListName].length === 1 && goodsInfo[specListName][0].list.length === 1 && goodsInfo[specListName][0].name === defaultSingleSkuName) {
        isMany = false;
      }
      return isMany;
    },
    // 默认价格区间计算
    priceCom() {
      let str = "";
      let that = this;
      let { selectShop = {}, goodsInfo = {}, skuListName, skuIdName } = that;
      if (selectShop[skuIdName]) {
        str = that.priceFilter(selectShop.price);
      } else {
        let skuList = goodsInfo[skuListName];
        if (skuList && skuList.length > 0) {
          let valueArr = [];
          skuList.map((skuItem, index) => {
            valueArr.push(skuItem.price);
          });
          let min = that.priceFilter(Math.min(...valueArr));
          let max = that.priceFilter(Math.max(...valueArr));
          if (min === max) {
            str = min + "";
          } else {
            str = `${min} - ${max}`;
          }
        }
      }
      return str;
    },
    // 库存显示
    stockCom() {
      let str = "";
      let that = this;
      let { selectShop = {}, goodsInfo = {}, skuListName, stockName } = that;
      if (selectShop[stockName]) {
        str = selectShop[stockName];
      } else {
        let skuList = goodsInfo[skuListName];
        if (skuList && skuList.length > 0) {
          let valueArr = [];
          skuList.map((skuItem, index) => {
            valueArr.push(skuItem[stockName]);
          });
          let min = Math.min(...valueArr);
          let max = Math.max(...valueArr);
          if (min === max) {
            str = min;
          } else {
            str = `${min} - ${max}`;
          }
        }
      }
      return str;
    }
  },
  watch: {
    valueCom(newVal, oldValue) {
      let that = this;
      if (newVal) {
        that.open();
      }
    },
    defaultGoods: {
      immediate: true,
      handler: function(newVal, oldValue) {
        let that = this;
        let { goodsIdName } = that;
        if (typeof newVal === "object" && newVal && newVal[goodsIdName] && !goodsCache[newVal[goodsIdName]]) {
          that.pushGoodsCache(newVal);
        }
      }
    }
  }
};
if (!Array) {
  const _easycom_vk_data_input_number_box2 = common_vendor.resolveComponent("vk-data-input-number-box");
  _easycom_vk_data_input_number_box2();
}
const _easycom_vk_data_input_number_box = () => "../vk-data-input-number-box/vk-data-input-number-box.js";
if (!Math) {
  _easycom_vk_data_input_number_box();
}
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return common_vendor.e({
    a: common_vendor.o(($event) => $options.close("mask")),
    b: $data.selectShop.image ? $data.selectShop.image : $data.goodsInfo[$props.goodsThumbName],
    c: $props.goodsThumbBackgroundColor,
    d: common_vendor.o((...args) => $options.previewImage && $options.previewImage(...args)),
    e: common_vendor.t($options.priceCom),
    f: common_vendor.n($options.priceCom.length > 16 ? "price2" : ""),
    g: $options.themeColorFn("priceColor"),
    h: !$props.hideStock
  }, !$props.hideStock ? {
    i: common_vendor.t($props.stockText),
    j: common_vendor.t($options.stockCom)
  } : {}, {
    k: common_vendor.t($data.selectArr.join(" ")),
    l: $options.isManyCom,
    m: common_vendor.f($data.goodsInfo[$props.specListName], (item, index1, i0) => {
      return {
        a: common_vendor.t(item.name),
        b: common_vendor.f(item.list, (item_value, index2, i1) => {
          return {
            a: common_vendor.t(item_value.name),
            b: index2,
            c: common_vendor.n(item_value.ishow ? "" : "noactived"),
            d: common_vendor.n($data.subIndex[index1] == index2 ? "actived" : ""),
            e: common_vendor.s(item_value.ishow ? "" : $options.themeColorFn("disableStyle")),
            f: common_vendor.s(item_value.ishow ? $options.themeColorFn("btnStyle") : ""),
            g: common_vendor.s($data.subIndex[index1] == index2 ? $options.themeColorFn("activedStyle") : ""),
            h: common_vendor.o(($event) => $options.skuClick(item_value, index1, index2), index2)
          };
        }),
        c: index1
      };
    }),
    n: $options.isManyCom,
    o: common_vendor.o($options.numChange),
    p: common_vendor.o(($event) => $data.selectNum = $event),
    q: common_vendor.p({
      min: $props.minBuyNum || 1,
      max: $options.maxBuyNumCom,
      step: $props.stepBuyNum || 1,
      ["step-strictly"]: $props.stepStrictly,
      ["positive-integer"]: true,
      modelValue: $data.selectNum
    }),
    r: $props.showClose != false
  }, $props.showClose != false ? {
    s: $props.closeImage,
    t: common_vendor.o(($event) => $options.close("close"))
  } : {}, {
    v: $data.outFoStock || $props.mode == 4
  }, $data.outFoStock || $props.mode == 4 ? {
    w: common_vendor.t($props.noStockText)
  } : $props.mode == 1 ? {
    y: common_vendor.t($props.addCartText),
    z: $options.themeColorFn("addCartColor"),
    A: $options.themeColorFn("addCartBackgroundColor"),
    B: common_vendor.o((...args) => $options.addCart && $options.addCart(...args)),
    C: common_vendor.t($props.buyNowText),
    D: $options.themeColorFn("buyNowColor"),
    E: $options.themeColorFn("buyNowBackgroundColor"),
    F: common_vendor.o((...args) => $options.buyNow && $options.buyNow(...args))
  } : $props.mode == 2 ? {
    H: common_vendor.t($props.addCartText),
    I: $options.themeColorFn("addCartColor"),
    J: $options.themeColorFn("addCartBackgroundColor"),
    K: common_vendor.o((...args) => $options.addCart && $options.addCart(...args))
  } : $props.mode == 3 ? {
    M: common_vendor.t($props.buyNowText),
    N: $options.themeColorFn("buyNowColor"),
    O: $options.themeColorFn("buyNowBackgroundColor"),
    P: common_vendor.o((...args) => $options.buyNow && $options.buyNow(...args))
  } : {}, {
    x: $props.mode == 1,
    G: $props.mode == 2,
    L: $props.mode == 3,
    Q: $props.safeAreaInsetBottom ? 1 : "",
    R: $props.borderRadius + "rpx " + $props.borderRadius + "rpx 0 0",
    S: $data.safeBottom + "px",
    T: common_vendor.n($options.valueCom && $data.complete ? "show" : "none"),
    U: common_vendor.o((...args) => $options.moveHandle && $options.moveHandle(...args)),
    V: common_vendor.o((...args) => $options.stop && $options.stop(...args))
  });
}
const Component = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__scopeId", "data-v-d1e15e37"], ["__file", "/Users/gygy/Desktop/毕业设计/SimplePick-Mall-UniUI/src/components/vk-data-goods-sku-popup/vk-data-goods-sku-popup.vue"]]);
wx.createComponent(Component);
//# sourceMappingURL=vk-data-goods-sku-popup.js.map
