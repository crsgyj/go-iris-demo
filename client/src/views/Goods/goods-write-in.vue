<template>
  <div class="container">
    <div class="limit-width center scroll-x">
      <!-- upload-input -->
      <input
        type="file"
        ref="file"
        style="visibility: hidden; width:0; height: 0"
        @change="onFileChange"
        required="application/msexcel"
        accept=".csv, application/vnd.ms-excel, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
      >
      <!-- upload-input:end -->
      <!-- btns -->
      <div class="btns">
        <el-button
          :type="list && list.length ? 'danger' : 'success' "
          @click="uploadTable"
        >{{list && list.length ? '删除表格' : '上传表格'}}</el-button>
        <el-button type="primary" @click="translate" :disabled="!list || !list.length">百度翻译</el-button>
        <el-button type="primary" :disabled="!canWriteIn" @click="writeIn">录入数据</el-button>
      </div>
      <!-- btns:end -->
      <!-- table -->
      <div ref="table">
        <el-table v-loading="loading" :data="list" border fit show-header>
          <el-table-column
            v-for="collumn in collumns"
            :key="collumn.prop"
            :prop="collumn.prop"
            :label="collumn.label"
            :width="collumn.width"
          ></el-table-column>
          <el-table-column label="操作">
            <template slot-scope="scope">
              <el-button size="mini" @click="handleEdit(scope.row)">修改</el-button>
              <el-popover placement="top" width="280">
                <p>确定删除"{{scope.row.name}}({{scope.row.goods_id}})"吗？</p>
                <div style="text-align: right; margin: 0">
                  <el-button size="mini" type="text" @click="$refs.table.click()">取消</el-button>
                  <el-button type="primary" size="mini" @click="handleDelete(scope.$index)">确定</el-button>
                </div>
                <el-button slot="reference" size="mini" type="danger">删除</el-button>
              </el-popover>
              <!-- <el-button size="mini" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button> -->
            </template>
          </el-table-column>
        </el-table>
      </div>
      <!-- table:end -->
      <!-- dialogs -->
      <edit-form @on-update-item="onUpdateItem" ref="editForm"></edit-form>
      <!-- dialogs:end -->
    </div>
  </div>
</template>

<script>
import XLSX from "xlsx";
import axios from "axios";
import axiosjsonp from "axios-jsonp";
import _ from "lodash";

import editForm from "./edit-form";

export default {
  components: {
    "edit-form": editForm
  },
  data() {
    return {
      list: null,
      workBook: null,
      hasTranslated: false,
      loading: false,
      collumns: [
        {
          prop: "goods_id",
          label: "货品编码",
          width: "160"
        },
        {
          prop: "name",
          label: "名称",
          width: "200"
        },
        {
          prop: "name_e",
          label: "名称(英文)",
          width: "200"
        },
        {
          prop: "addr",
          label: "地址",
          width: "300"
        },
        {
          prop: "addr_e",
          label: "地址(英文)",
          width: "300"
        }
      ]
    };
  },
  computed: {
    canWriteIn() {
      return this.list && this.list.length && this.hasTranslated;
    }
  },
  watch: {},
  methods: {
    // 上传表格
    uploadTable() {
      this.hasTranslated = false;
      if (this.list && this.list.length) {
        return this.deleteTable();
      }
      // 清除上次选择
      this.$refs.file.value = "";
      this.$refs.file.click();
    },
    // 文件被上传
    onFileChange() {
      this.readExcel();
    },
    // 读取Excel -> workbook
    readExcel() {
      let self = this;
      let file = this.$refs.file.files[0];
      if (!file) return;
      let reader = new FileReader();
      reader.readAsArrayBuffer(file);
      reader.onload = function(e) {
        let workBook = XLSX.read(e.target.result, { type: "buffer" });
        if (!workBook.Workbook) {
          return;
        }
        self.workBook = workBook;

        let firstSheetName = workBook.SheetNames[0];
        let firstSheet = workBook.Sheets[firstSheetName];
        let goodsData = XLSX.utils.sheet_to_json(firstSheet, {
          header: "A",
          blankrows: false,
          defval: "",
          // range: "" range可以用户选择
        });
        goodsData = goodsData.map(e => ({
          goods_id: String(e.A),
          name: String(e.B),
          name_e: "",
          addr: String(e.C),
          addr_e: ""
        }));
        self.list = goodsData;
      };
    },
    deleteTable() {
      this.list = null;
    },
    async translate() {
      this.hasTranslated = false;
      let chunckList = _.chunk(this.list, 100);
      let result = [];
      this.loading = true;
      for (let i = 0, len = chunckList.length; i < len; i++) {
        let chunck = chunckList[i];
        try {
          let chunckRst = await this.commTask(chunck);
          result = result.concat(chunckRst);
        } catch (err) {
          let errResp = err.response;

          if (errResp && errResp.status === 401) {
            this.$emit("not-login");
          } else {
            this.$message({
              message: "翻译失败。" + err.message,
              type: "error"
            });
          }

          this.loading = false;

          return;
        }
      }
      this.list = result;
      this.hasTranslated = true;
      this.loading = false;
      this.$message({
        message: "翻译完成。",
        type: "success"
      });
    },
    async commTask(chunck) {
      let access = {};
      let q = encodeURIComponent(
        chunck.map(e => `${e.name || "-"}\n${e.addr || "-"}`).join("\n")
      );
      let { resp, err } = await this.fetchBaiduAccess(q);
      if (err) {
        throw err;
      } 
      access = resp.data;
      try {
        let resp = await this.$http.get(this.$config.baiduFY, {
          params: {
            q: access.q,
            appid: access.appid,
            salt: access.salt,
            from: access.from,
            to: access.to,
            sign: access.sign
          },
          adapter: axiosjsonp
        });
        let trans_result = resp.data.trans_result;
        for (let i in chunck) {
          chunck[i].name_e = trans_result[i * 2].dst;
          chunck[i].addr_e = trans_result[i * 2 + 1].dst;
        }
      } catch (err) {
        throw new Error("baiduFY error: ", err.message);
      }

      return chunck
    },
    writeIn() {
      if (!this.canWriteIn) {
        return;
      }
      this.loading = true;
      this.$api
        .post("/goods/list", {
          list: this.list
        })
        .then(res => {
          this.loading = false;
          this.$message({
            message: "录入成功。",
            type: "success"
          });
        })
        .catch(err => {
          this.loading = false;
          let resp = err.response;
          if (resp && resp.status === 401) {
            this.$emit("not-login");
          }
          this.$message({
            message: "录入失败。" + resp.response.data.message,
            type: "error"
          });
          return;
        });
    },
    // 获取百度翻译token
    fetchBaiduAccess(q) {
      return new Promise((resolve, reject) => {
        this.$api
          .get("/baiduFY/access?q=" + q)
          .then(resp => {
            resolve({ resp: resp, err: null });
          })
          .catch(err => {
            resolve({ resp: null, err: err });
          });
      });
    },
    // 删除单条
    handleDelete(i) {
      this.list.splice(i, 1);
      this.$message("已删除。");
      if (!this.list.length) {
        this.list = null;
        this.hasTranslated = false;
      }
      // 清除dialog
      this.$refs.table.click();
    },
    // 修改
    handleEdit(row) {
      this.$refs.editForm.open(row);
    },
    // 更新单条
    onUpdateItem(data, cb) {
      if (!data) return cb();
      let list = this.list.slice();
      for (let i = 0, len = list.length; i < len; i++) {
        if (list[i].goods_id === data.goods_id) {
          list.splice(i, 1, data);
          break;
        }
      }
      this.list = list;

      cb();
    }
  }
};
</script>

<style lang="stylus" scoped>
.container
  width 100%
  box-sizing border-box
  height 100vh
  background #fefefe
  padding 16px
  .btns
    text-align left
    margin-bottom 16px
</style>
