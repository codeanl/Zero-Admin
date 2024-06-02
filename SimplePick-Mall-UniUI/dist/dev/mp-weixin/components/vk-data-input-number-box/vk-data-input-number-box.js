"use strict";
const common_vendor = require("../../common/vendor.js");
const _sfc_main = {
  name: "vk-data-input-number-box",
  emits: ["update:modelValue", "input", "change", "blur", "plus", "minus"],
  props: {
    // 预显示的数字
    value: {
      type: Number,
      default: 1
    },
    modelValue: {
      type: Number,
      default: 1
    },
    // 背景颜色
    bgColor: {
      type: String,
      default: "#F2F3F5"
    },
    // 最小值
    min: {
      type: Number,
      default: 0
    },
    // 最大值
    max: {
      type: Number,
      default: 99999
    },
    // 步进值，每次加或减的值
    step: {
      type: Number,
      default: 1
    },
    // 步进值，首次增加或最后减的值
    stepFirst: {
      type: Number,
      default: 0
    },
    // 是否只能输入 step 的倍数
    stepStrictly: {
      type: Boolean,
      default: false
    },
    // 是否禁用加减操作
    disabled: {
      type: Boolean,
      default: false
    },
    // input的字体大小，单位rpx
    size: {
      type: [Number, String],
      default: 26
    },
    // 加减图标的颜色
    color: {
      type: String,
      default: "#323233"
    },
    // input宽度，单位rpx
    inputWidth: {
      type: [Number, String],
      default: 80
    },
    // input高度，单位rpx
    inputHeight: {
      type: [Number, String],
      default: 50
    },
    // index索引，用于列表中使用，让用户知道是哪个numberbox发生了变化，一般使用for循环出来的index值即可
    index: {
      type: [Number, String],
      default: ""
    },
    // 是否禁用输入框，与disabled作用于输入框时，为OR的关系，即想要禁用输入框，又可以加减的话
    // 设置disabled为false，disabledInput为true即可
    disabledInput: {
      type: Boolean,
      default: false
    },
    // 输入框于键盘之间的距离
    cursorSpacing: {
      type: [Number, String],
      default: 100
    },
    // 是否开启长按连续递增或递减
    longPress: {
      type: Boolean,
      default: true
    },
    // 开启长按触发后，每触发一次需要多久
    pressTime: {
      type: [Number, String],
      default: 250
    },
    // 是否只能输入大于或等于0的整数(正整数)
    positiveInteger: {
      type: Boolean,
      default: true
    }
  },
  watch: {
    valueCom(v1, v2) {
      if (!this.changeFromInner) {
        this.inputVal = v1;
        this.$nextTick(function() {
          this.changeFromInner = false;
        });
      }
    },
    inputVal(v1, v2) {
      if (v1 == "")
        return;
      let value = 0;
      let tmp = this.isNumber(v1);
      if (tmp && v1 >= this.min && v1 <= this.max)
        value = v1;
      else
        value = v2;
      if (this.positiveInteger) {
        if (v1 < 0 || String(v1).indexOf(".") !== -1) {
          value = v2;
          this.$nextTick(() => {
            this.inputVal = v2;
          });
        }
      }
      this.handleChange(value, "change");
    },
    min(v1) {
      if (v1 !== void 0 && v1 != "" && this.valueCom < v1) {
        this.$emit("input", v1);
        this.$emit("update:modelValue", v1);
      }
    },
    max(v1) {
      if (v1 !== void 0 && v1 != "" && this.valueCom > v1) {
        this.$emit("input", v1);
        this.$emit("update:modelValue", v1);
      }
    }
  },
  data() {
    return {
      inputVal: 1,
      // 输入框中的值，不能直接使用props中的value，因为应该改变props的状态
      timer: null,
      // 用作长按的定时器
      changeFromInner: false,
      // 值发生变化，是来自内部还是外部
      innerChangeTimer: null
      // 内部定时器
    };
  },
  created() {
    this.inputVal = Number(this.valueCom);
  },
  computed: {
    valueCom() {
      return this.modelValue;
    },
    getCursorSpacing() {
      return Number(common_vendor.index.upx2px(this.cursorSpacing));
    }
  },
  methods: {
    // 点击退格键
    btnTouchStart(callback) {
      this[callback]();
      if (!this.longPress)
        return;
      clearInterval(this.timer);
      this.timer = null;
      this.timer = setInterval(() => {
        this[callback]();
      }, this.pressTime);
    },
    clearTimer() {
      this.$nextTick(() => {
        clearInterval(this.timer);
        this.timer = null;
      });
    },
    minus() {
      this.computeVal("minus");
    },
    plus() {
      this.computeVal("plus");
    },
    // 为了保证小数相加减出现精度溢出的问题
    calcPlus(num1, num2) {
      let baseNum, baseNum1, baseNum2;
      try {
        baseNum1 = num1.toString().split(".")[1].length;
      } catch (e) {
        baseNum1 = 0;
      }
      try {
        baseNum2 = num2.toString().split(".")[1].length;
      } catch (e) {
        baseNum2 = 0;
      }
      baseNum = Math.pow(10, Math.max(baseNum1, baseNum2));
      let precision = baseNum1 >= baseNum2 ? baseNum1 : baseNum2;
      return ((num1 * baseNum + num2 * baseNum) / baseNum).toFixed(precision);
    },
    // 为了保证小数相加减出现精度溢出的问题
    calcMinus(num1, num2) {
      let baseNum, baseNum1, baseNum2;
      try {
        baseNum1 = num1.toString().split(".")[1].length;
      } catch (e) {
        baseNum1 = 0;
      }
      try {
        baseNum2 = num2.toString().split(".")[1].length;
      } catch (e) {
        baseNum2 = 0;
      }
      baseNum = Math.pow(10, Math.max(baseNum1, baseNum2));
      let precision = baseNum1 >= baseNum2 ? baseNum1 : baseNum2;
      return ((num1 * baseNum - num2 * baseNum) / baseNum).toFixed(precision);
    },
    computeVal(type) {
      common_vendor.index.hideKeyboard();
      if (this.disabled)
        return;
      let value = 0;
      if (type === "minus") {
        if (this.stepFirst > 0 && this.inputVal == this.stepFirst) {
          value = this.min;
        } else {
          value = this.calcMinus(this.inputVal, this.step);
        }
      } else if (type === "plus") {
        if (this.stepFirst > 0 && this.inputVal < this.stepFirst) {
          value = this.stepFirst;
        } else {
          value = this.calcPlus(this.inputVal, this.step);
        }
      }
      if (this.stepStrictly) {
        let strictly = value % this.step;
        if (strictly > 0) {
          value -= strictly;
        }
      }
      if (value > this.max) {
        value = this.max;
      } else if (value < this.min) {
        value = this.min;
      }
      this.inputVal = value;
      this.handleChange(value, type);
    },
    // 处理用户手动输入的情况
    onBlur(event) {
      let val = 0;
      let value = event.detail.value;
      if (!/(^\d+$)/.test(value) || value[0] == 0)
        val = this.min;
      val = +value;
      if (this.stepFirst > 0 && this.inputVal < this.stepFirst && this.inputVal > 0) {
        val = this.stepFirst;
      }
      if (this.stepStrictly) {
        let strictly = val % this.step;
        if (strictly > 0) {
          val -= strictly;
        }
      }
      if (val > this.max) {
        val = this.max;
      } else if (val < this.min) {
        val = this.min;
      }
      this.$nextTick(() => {
        this.inputVal = val;
      });
      this.handleChange(val, "blur");
    },
    handleChange(value, type) {
      if (this.disabled)
        return;
      if (this.innerChangeTimer) {
        clearTimeout(this.innerChangeTimer);
        this.innerChangeTimer = null;
      }
      this.changeFromInner = true;
      this.innerChangeTimer = setTimeout(() => {
        this.changeFromInner = false;
      }, 150);
      this.$emit("input", Number(value));
      this.$emit("update:modelValue", Number(value));
      this.$emit(type, {
        // 转为Number类型
        value: Number(value),
        index: this.index
      });
    },
    /**
     * 验证十进制数字
     */
    isNumber(value) {
      return /^(?:-?\d+|-?\d{1,3}(?:,\d{3})+)?(?:\.\d+)?$/.test(value);
    }
  }
};
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return {
    a: common_vendor.s("font-size:" + (Number($props.size) + 10) + "rpx"),
    b: $props.disabled || $data.inputVal <= $props.min ? 1 : "",
    c: $props.bgColor,
    d: $props.inputHeight + "rpx",
    e: $props.color,
    f: $props.size + "rpx",
    g: common_vendor.o(($event) => $options.btnTouchStart("minus")),
    h: common_vendor.o((...args) => $options.clearTimer && $options.clearTimer(...args)),
    i: $props.disabledInput || $props.disabled,
    j: $options.getCursorSpacing,
    k: $props.disabled ? 1 : "",
    l: $props.color,
    m: $props.size + "rpx",
    n: $props.bgColor,
    o: $props.inputHeight + "rpx",
    p: $props.inputWidth + "rpx",
    q: common_vendor.o((...args) => $options.onBlur && $options.onBlur(...args)),
    r: $data.inputVal,
    s: common_vendor.o(($event) => $data.inputVal = $event.detail.value),
    t: common_vendor.s("font-size:" + (Number($props.size) + 10) + "rpx"),
    v: $props.disabled || $data.inputVal >= $props.max ? 1 : "",
    w: $props.bgColor,
    x: $props.inputHeight + "rpx",
    y: $props.color,
    z: $props.size + "rpx",
    A: common_vendor.o(($event) => $options.btnTouchStart("plus")),
    B: common_vendor.o((...args) => $options.clearTimer && $options.clearTimer(...args))
  };
}
const Component = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__scopeId", "data-v-3517706a"], ["__file", "/Users/gygy/Desktop/毕业设计/SimplePick-Mall-UniUI/src/components/vk-data-input-number-box/vk-data-input-number-box.vue"]]);
wx.createComponent(Component);
//# sourceMappingURL=vk-data-input-number-box.js.map
