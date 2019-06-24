// modules are defined as an array
// [ module function, map of requires ]
//
// map of requires is short require name -> numeric require
//
// anything defined in a previous bundle is accessed via the
// orig method which is the require for previous bundles
parcelRequire = (function (modules, cache, entry, globalName) {
  // Save the require from previous bundle to this closure if any
  var previousRequire = typeof parcelRequire === 'function' && parcelRequire;
  var nodeRequire = typeof require === 'function' && require;

  function newRequire(name, jumped) {
    if (!cache[name]) {
      if (!modules[name]) {
        // if we cannot find the module within our internal map or
        // cache jump to the current global require ie. the last bundle
        // that was added to the page.
        var currentRequire = typeof parcelRequire === 'function' && parcelRequire;
        if (!jumped && currentRequire) {
          return currentRequire(name, true);
        }

        // If there are other bundles on this page the require from the
        // previous one is saved to 'previousRequire'. Repeat this as
        // many times as there are bundles until the module is found or
        // we exhaust the require chain.
        if (previousRequire) {
          return previousRequire(name, true);
        }

        // Try the node require function if it exists.
        if (nodeRequire && typeof name === 'string') {
          return nodeRequire(name);
        }

        var err = new Error('Cannot find module \'' + name + '\'');
        err.code = 'MODULE_NOT_FOUND';
        throw err;
      }

      localRequire.resolve = resolve;
      localRequire.cache = {};

      var module = cache[name] = new newRequire.Module(name);

      modules[name][0].call(module.exports, localRequire, module, module.exports, this);
    }

    return cache[name].exports;

    function localRequire(x){
      return newRequire(localRequire.resolve(x));
    }

    function resolve(x){
      return modules[name][1][x] || x;
    }
  }

  function Module(moduleName) {
    this.id = moduleName;
    this.bundle = newRequire;
    this.exports = {};
  }

  newRequire.isParcelRequire = true;
  newRequire.Module = Module;
  newRequire.modules = modules;
  newRequire.cache = cache;
  newRequire.parent = previousRequire;
  newRequire.register = function (id, exports) {
    modules[id] = [function (require, module) {
      module.exports = exports;
    }, {}];
  };

  var error;
  for (var i = 0; i < entry.length; i++) {
    try {
      newRequire(entry[i]);
    } catch (e) {
      // Save first error but execute all entries
      if (!error) {
        error = e;
      }
    }
  }

  if (entry.length) {
    // Expose entry point to Node, AMD or browser globals
    // Based on https://github.com/ForbesLindesay/umd/blob/master/template.js
    var mainExports = newRequire(entry[entry.length - 1]);

    // CommonJS
    if (typeof exports === "object" && typeof module !== "undefined") {
      module.exports = mainExports;

    // RequireJS
    } else if (typeof define === "function" && define.amd) {
     define(function () {
       return mainExports;
     });

    // <script>
    } else if (globalName) {
      this[globalName] = mainExports;
    }
  }

  // Override the current require with this new one
  parcelRequire = newRequire;

  if (error) {
    // throw error from earlier, _after updating parcelRequire_
    throw error;
  }

  return newRequire;
})({"src/views/Goods/goods-list.vue":[function(require,module,exports) {
"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.default = void 0;

var _editForm = _interopRequireDefault(require("./editForm"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
var _default = {
  components: {
    "edit-form": _editForm.default
  },
  data: function data() {
    return {
      querys: {
        goodsID: "",
        perPage: 20,
        page: 1
      },
      loading: false,
      listData: {
        data: [],
        count: 0
      },
      collumns: [{
        prop: "goods_id",
        label: "货品编码",
        width: "160"
      }, {
        prop: "name",
        label: "名称",
        width: "200"
      }, {
        prop: "name_e",
        label: "名称(英文)",
        width: "200"
      }, {
        prop: "addr",
        label: "地址",
        width: "300"
      }, {
        prop: "addr_e",
        label: "地址(英文)",
        width: "300"
      }],
      timeout: null
    };
  },
  created: function created() {
    this.search();
  },
  activated: function activated() {
    this.search();
  },
  methods: {
    search: function search(page) {
      var _this = this;

      if (!page) {
        this.querys.page = 1;
      } else {
        this.querys.page = page;
      }

      var specGoodsID = this.querys.goodsID;
      this.loading = true;
      this.$api.get("/goods/list", {
        params: {
          page: this.querys.page,
          per_page: this.querys.perPage,
          goods_id: this.querys.goodsID
        }
      }).then(function (res) {
        _this.listData.data = res.data.list;
        _this.listData.count = res.data.count;
        _this.loading = false;
      }).catch(function (err) {
        _this.loading = false;
        var resp = err.response;

        if (!resp) {
          return alert(resp.message);
        }

        if (resp.status === 401) {
          _this.$emit("not-login");
        }
      });
    },
    // 输入建议
    suggest: function suggest(goods_id, cb) {
      var _this2 = this;

      if (!goods_id) {
        return cb(null);
      }

      clearTimeout(this.timeout);
      this.$api.get("/goods/suggest", {
        params: {
          goods_id: goods_id
        }
      }).then(function (res) {
        _this2.timeout = setTimeout(function () {
          cb(res.data.map(function (e) {
            return {
              value: e
            };
          }));
        }, 500);
      }).catch(function (err) {});
    },
    // 删除
    handleDelete: function handleDelete(row) {
      var _this3 = this;

      var goodsId = row.goods_id;
      this.loading = true;
      this.$api.delete("goods/".concat(goodsId)).then(function (res) {
        _this3.$message("已删除。");

        _this3.search(_this3.querys.page);

        _this3.$refs.table.click();

        _this3.loading = false;
      }).catch(function (err) {
        _this3.loading = false;
        var resp = err.response;

        if (!resp) {
          return _this3.$message.error(err.message);
        } // 未登录


        if (resp.status === 401) {
          return _this3.$emit("not-login");
        } // 具体报错


        var reason = function () {
          var str = "";

          try {
            str = resp.data.message;
          } catch (e) {
            console.log(e);
          }

          return str;
        }();

        _this3.$message.error("删除失败。" + reason);
      });
    },
    // 修改
    handleEdit: function handleEdit(row) {
      this.$refs.editForm.open(row);
    },
    onUpdateItem: function onUpdateItem(data, cb) {
      var _this4 = this;

      if (!data) return cb();
      this.loading = true;
      var action = this.$api.put("/goods/".concat(data.goods_id), {
        item: data
      }).then(function () {
        _this4.search(_this4.querys.page);
      });
      this.loading = false;
      cb(action);
    }
  }
};
exports.default = _default;
        var $de0b55 = exports.default || module.exports;
      
      if (typeof $de0b55 === 'function') {
        $de0b55 = $de0b55.options;
      }
    
        /* template */
        Object.assign($de0b55, (function () {
          var render = function() {
  var _vm = this
  var _h = _vm.$createElement
  var _c = _vm._self._c || _h
  return _c("div", { staticClass: "container" }, [
    _c(
      "div",
      { staticClass: "limit-width center scroll-x" },
      [
        _c(
          "div",
          {
            staticClass: "search-box",
            on: {
              keydown: function($event) {
                if (
                  !$event.type.indexOf("key") &&
                  _vm._k($event.keyCode, "enter", 13, $event.key, "Enter")
                ) {
                  return null
                }
                return _vm.search()
              }
            }
          },
          [
            _c("el-autocomplete", {
              staticClass: "inline-input",
              attrs: {
                "fetch-suggestions": _vm.suggest,
                placeholder: "请输入货品编码",
                "trigger-on-focus": false
              },
              model: {
                value: _vm.querys.goodsID,
                callback: function($$v) {
                  _vm.$set(_vm.querys, "goodsID", $$v)
                },
                expression: "querys.goodsID"
              }
            }),
            _vm._v(" "),
            _c(
              "el-button",
              {
                attrs: { type: "primary" },
                on: {
                  click: function($event) {
                    return _vm.search()
                  }
                }
              },
              [_vm._v("查找")]
            )
          ],
          1
        ),
        _vm._v(" "),
        _c(
          "div",
          { ref: "table", staticStyle: { "min-height": "500px" } },
          [
            _c(
              "el-table",
              {
                directives: [
                  {
                    name: "loading",
                    rawName: "v-loading",
                    value: _vm.loading,
                    expression: "loading"
                  }
                ],
                attrs: {
                  data: _vm.listData.data,
                  border: "",
                  fit: "",
                  "show-header": ""
                }
              },
              [
                _vm._l(_vm.collumns, function(collumn) {
                  return _c("el-table-column", {
                    key: collumn.prop,
                    attrs: {
                      prop: collumn.prop,
                      label: collumn.label,
                      width: collumn.width
                    }
                  })
                }),
                _vm._v(" "),
                _c("el-table-column", {
                  attrs: { label: "操作" },
                  scopedSlots: _vm._u([
                    {
                      key: "default",
                      fn: function(scope) {
                        return [
                          _c(
                            "el-button",
                            {
                              attrs: { size: "mini" },
                              on: {
                                click: function($event) {
                                  return _vm.handleEdit(scope.row)
                                }
                              }
                            },
                            [_vm._v("修改")]
                          ),
                          _vm._v(" "),
                          _c(
                            "el-popover",
                            { attrs: { placement: "top", width: "280" } },
                            [
                              _c("p", [
                                _vm._v(
                                  '确定删除"' +
                                    _vm._s(scope.row.name) +
                                    "(" +
                                    _vm._s(scope.row.goods_id) +
                                    ')"吗？'
                                )
                              ]),
                              _vm._v(" "),
                              _c(
                                "div",
                                {
                                  staticStyle: {
                                    "text-align": "right",
                                    margin: "0"
                                  }
                                },
                                [
                                  _c(
                                    "el-button",
                                    {
                                      attrs: { size: "mini", type: "text" },
                                      on: {
                                        click: function($event) {
                                          return _vm.$refs.table.click()
                                        }
                                      }
                                    },
                                    [_vm._v("取消")]
                                  ),
                                  _vm._v(" "),
                                  _c(
                                    "el-button",
                                    {
                                      attrs: { type: "primary", size: "mini" },
                                      on: {
                                        click: function($event) {
                                          return _vm.handleDelete(scope.row)
                                        }
                                      }
                                    },
                                    [_vm._v("确定")]
                                  )
                                ],
                                1
                              ),
                              _vm._v(" "),
                              _c(
                                "el-button",
                                {
                                  attrs: {
                                    slot: "reference",
                                    size: "mini",
                                    type: "danger"
                                  },
                                  slot: "reference"
                                },
                                [_vm._v("删除")]
                              )
                            ],
                            1
                          )
                        ]
                      }
                    }
                  ])
                })
              ],
              2
            )
          ],
          1
        ),
        _vm._v(" "),
        _c(
          "div",
          { staticClass: "text-align-right" },
          [
            _c("el-pagination", {
              attrs: {
                "page-size": _vm.querys.perPage,
                "pager-count": 11,
                layout: "total, prev, pager, next",
                total: _vm.listData.count
              }
            })
          ],
          1
        ),
        _vm._v(" "),
        _c("edit-form", {
          ref: "editForm",
          on: { "on-update-item": _vm.onUpdateItem }
        })
      ],
      1
    )
  ])
}
var staticRenderFns = []
render._withStripped = true

          return {
            render: render,
            staticRenderFns: staticRenderFns,
            _compiled: true,
            _scopeId: "data-v-de0b55",
            functional: undefined
          };
        })());
      
    /* hot reload */
    (function () {
      if (module.hot) {
        var api = require('vue-hot-reload-api');
        api.install(require('vue'));
        if (api.compatible) {
          module.hot.accept();
          if (!module.hot.data) {
            api.createRecord('$de0b55', $de0b55);
          } else {
            api.reload('$de0b55', $de0b55);
          }
        }

        
        var reloadCSS = require('_css_loader');
        module.hot.dispose(reloadCSS);
        module.hot.accept(reloadCSS);
      
      }
    })();
},{"./editForm":"src/views/Goods/editForm.vue","_css_loader":"node_modules/parcel-bundler/src/builtins/css-loader.js","vue-hot-reload-api":"node_modules/vue-hot-reload-api/dist/index.js","vue":"node_modules/vue/dist/vue.common.js"}],"node_modules/parcel-bundler/src/builtins/hmr-runtime.js":[function(require,module,exports) {
var global = arguments[3];
var OVERLAY_ID = '__parcel__error__overlay__';
var OldModule = module.bundle.Module;

function Module(moduleName) {
  OldModule.call(this, moduleName);
  this.hot = {
    data: module.bundle.hotData,
    _acceptCallbacks: [],
    _disposeCallbacks: [],
    accept: function (fn) {
      this._acceptCallbacks.push(fn || function () {});
    },
    dispose: function (fn) {
      this._disposeCallbacks.push(fn);
    }
  };
  module.bundle.hotData = null;
}

module.bundle.Module = Module;
var checkedAssets, assetsToAccept;
var parent = module.bundle.parent;

if ((!parent || !parent.isParcelRequire) && typeof WebSocket !== 'undefined') {
  var hostname = "" || location.hostname;
  var protocol = location.protocol === 'https:' ? 'wss' : 'ws';
  var ws = new WebSocket(protocol + '://' + hostname + ':' + "55536" + '/');

  ws.onmessage = function (event) {
    checkedAssets = {};
    assetsToAccept = [];
    var data = JSON.parse(event.data);

    if (data.type === 'update') {
      var handled = false;
      data.assets.forEach(function (asset) {
        if (!asset.isNew) {
          var didAccept = hmrAcceptCheck(global.parcelRequire, asset.id);

          if (didAccept) {
            handled = true;
          }
        }
      }); // Enable HMR for CSS by default.

      handled = handled || data.assets.every(function (asset) {
        return asset.type === 'css' && asset.generated.js;
      });

      if (handled) {
        console.clear();
        data.assets.forEach(function (asset) {
          hmrApply(global.parcelRequire, asset);
        });
        assetsToAccept.forEach(function (v) {
          hmrAcceptRun(v[0], v[1]);
        });
      } else {
        window.location.reload();
      }
    }

    if (data.type === 'reload') {
      ws.close();

      ws.onclose = function () {
        location.reload();
      };
    }

    if (data.type === 'error-resolved') {
      console.log('[parcel] ✨ Error resolved');
      removeErrorOverlay();
    }

    if (data.type === 'error') {
      console.error('[parcel] 🚨  ' + data.error.message + '\n' + data.error.stack);
      removeErrorOverlay();
      var overlay = createErrorOverlay(data);
      document.body.appendChild(overlay);
    }
  };
}

function removeErrorOverlay() {
  var overlay = document.getElementById(OVERLAY_ID);

  if (overlay) {
    overlay.remove();
  }
}

function createErrorOverlay(data) {
  var overlay = document.createElement('div');
  overlay.id = OVERLAY_ID; // html encode message and stack trace

  var message = document.createElement('div');
  var stackTrace = document.createElement('pre');
  message.innerText = data.error.message;
  stackTrace.innerText = data.error.stack;
  overlay.innerHTML = '<div style="background: black; font-size: 16px; color: white; position: fixed; height: 100%; width: 100%; top: 0px; left: 0px; padding: 30px; opacity: 0.85; font-family: Menlo, Consolas, monospace; z-index: 9999;">' + '<span style="background: red; padding: 2px 4px; border-radius: 2px;">ERROR</span>' + '<span style="top: 2px; margin-left: 5px; position: relative;">🚨</span>' + '<div style="font-size: 18px; font-weight: bold; margin-top: 20px;">' + message.innerHTML + '</div>' + '<pre>' + stackTrace.innerHTML + '</pre>' + '</div>';
  return overlay;
}

function getParents(bundle, id) {
  var modules = bundle.modules;

  if (!modules) {
    return [];
  }

  var parents = [];
  var k, d, dep;

  for (k in modules) {
    for (d in modules[k][1]) {
      dep = modules[k][1][d];

      if (dep === id || Array.isArray(dep) && dep[dep.length - 1] === id) {
        parents.push(k);
      }
    }
  }

  if (bundle.parent) {
    parents = parents.concat(getParents(bundle.parent, id));
  }

  return parents;
}

function hmrApply(bundle, asset) {
  var modules = bundle.modules;

  if (!modules) {
    return;
  }

  if (modules[asset.id] || !bundle.parent) {
    var fn = new Function('require', 'module', 'exports', asset.generated.js);
    asset.isNew = !modules[asset.id];
    modules[asset.id] = [fn, asset.deps];
  } else if (bundle.parent) {
    hmrApply(bundle.parent, asset);
  }
}

function hmrAcceptCheck(bundle, id) {
  var modules = bundle.modules;

  if (!modules) {
    return;
  }

  if (!modules[id] && bundle.parent) {
    return hmrAcceptCheck(bundle.parent, id);
  }

  if (checkedAssets[id]) {
    return;
  }

  checkedAssets[id] = true;
  var cached = bundle.cache[id];
  assetsToAccept.push([bundle, id]);

  if (cached && cached.hot && cached.hot._acceptCallbacks.length) {
    return true;
  }

  return getParents(global.parcelRequire, id).some(function (id) {
    return hmrAcceptCheck(global.parcelRequire, id);
  });
}

function hmrAcceptRun(bundle, id) {
  var cached = bundle.cache[id];
  bundle.hotData = {};

  if (cached) {
    cached.hot.data = bundle.hotData;
  }

  if (cached && cached.hot && cached.hot._disposeCallbacks.length) {
    cached.hot._disposeCallbacks.forEach(function (cb) {
      cb(bundle.hotData);
    });
  }

  delete bundle.cache[id];
  bundle(id);
  cached = bundle.cache[id];

  if (cached && cached.hot && cached.hot._acceptCallbacks.length) {
    cached.hot._acceptCallbacks.forEach(function (cb) {
      cb();
    });

    return true;
  }
}
},{}]},{},["node_modules/parcel-bundler/src/builtins/hmr-runtime.js"], "app")
//# sourceMappingURL=/goods-list.143f889c.js.map