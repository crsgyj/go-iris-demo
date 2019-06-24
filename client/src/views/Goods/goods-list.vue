<template>
  <div class="container">
    <div class="limit-width center scroll-x">
      <!-- search -->
      <div class="search-box" @keydown.enter="search()">
        <el-autocomplete
          class="inline-input"
          v-model="querys.goodsID"
          :fetch-suggestions="suggest"
          placeholder="请输入货品编码"
          :trigger-on-focus="false"
        ></el-autocomplete>
        <el-button type="primary" @click="search()">查找</el-button>
      </div>
      <!-- search:end -->
      <!-- table -->
      <div style="min-height: 500px" ref="table">
        <el-table v-loading="loading" :data="listData.data" border fit show-header>
          <el-table-column
            v-for="collumn in collumns"
            :key="collumn.prop"
            :prop="collumn.prop"
            :label="collumn.label"
            :width="collumn.width"
            :formatter="collumn.formatter"
          ></el-table-column>
          <el-table-column label="操作">
            <template slot-scope="scope">
              <el-button size="mini" @click="handleEdit(scope.row)">修改</el-button>
              <el-popover placement="top" width="280">
                <p>确定删除"{{scope.row.name}}({{scope.row.goods_id}})"吗？</p>
                <div style="text-align: right; margin: 0">
                  <el-button size="mini" type="text" @click="$refs.table.click()">取消</el-button>
                  <el-button type="primary" size="mini" @click="handleDelete(scope.row)">确定</el-button>
                </div>
                <el-button slot="reference" size="mini" type="danger">删除</el-button>
              </el-popover>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <!-- table:end -->
      <!-- pager -->
      <div class="text-align-right">
        <el-pagination
          :page-size="querys.perPage"
          :pager-count="11"
          layout="total, prev, pager, next"
          :total="listData.count"
        ></el-pagination>
      </div>
      <!-- pager:end -->
      <!-- dialogs -->
      <edit-form @on-update-item="onUpdateItem" ref="editForm"></edit-form>
      <!-- dialogs:end -->
    </div>
  </div>
</template>

<script>
import editForm from "./edit-form";

export default {
  components: {
    "edit-form": editForm
  },
  data() {
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
        },
        {
          prop: "created_date",
          label: "录入时间",
          width: "120",
          formatter: ({ created_date }) => {
            return created_date.replace(/(\d{4})(\d{2})(\d{2})/g, '$1/$2/$3')
          }
        }
      ],
      timeout: null
    };
  },
  created() {
    this.search();
  },
  activated() {
    this.search();
  },
  methods: {
    search(page) {
      if (!page) {
        this.querys.page = 1;
      } else {
        this.querys.page = page;
      }
      let specGoodsID = this.querys.goodsID;
      this.loading = true;
      this.$api
        .get("/goods/list", {
          params: {
            page: this.querys.page,
            per_page: this.querys.perPage,
            goods_id: this.querys.goodsID
          }
        })
        .then(res => {
          this.listData.data = res.data.list;
          this.listData.count = res.data.count;
          this.loading = false;
        })
        .catch(err => {
          this.loading = false;
          let resp = err.response;
          if (!resp) {
            return alert(resp.message);
          }
          if (resp.status === 401) {
            this.$emit("not-login");
          }
        });
    },
    // 输入建议
    suggest(goods_id, cb) {
      if (!goods_id) {
        return cb(null);
      }
      clearTimeout(this.timeout);
      this.$api
        .get("/goods/suggest", {
          params: {
            goods_id: goods_id
          }
        })
        .then(res => {
          this.timeout = setTimeout(() => {
            cb(res.data.map(e => ({ value: e })));
          }, 500);
        })
        .catch(err => {});
    },
    // 删除
    handleDelete(row) {
      let goodsId = row.goods_id;
      this.loading = true;
      this.$api
        .delete(`goods/${goodsId}`)
        .then(res => {
          this.$message("已删除。");
          this.search(this.querys.page);
          this.$refs.table.click();
          this.loading = false;
        })
        .catch(err => {
          this.loading = false;
          let resp = err.response;
          if (!resp) {
            return this.$message.error(err.message);
          }
          // 未登录
          if (resp.status === 401) {
            return this.$emit("not-login");
          }
          // 具体报错
          let reason = (function() {
            let str = "";
            try {
              str = resp.data.message;
            } catch (e) {}
            return str;
          })();
          this.$message.error("删除失败。" + reason);
        });
    },
    // 修改
    handleEdit(row) {
      this.$refs.editForm.open(row);
    },
    onUpdateItem(data, cb) {
      if (!data) return cb();
      this.loading = true;

      let action = this.$api
        .put(`/goods/${data.goods_id}`, { item: data })
        .then(() => {
          this.search(this.querys.page);
        });
      this.loading = false;
      cb(action);
    }
  }
};
</script>

<style lang="stylus" scoped>
.container
  width 100%
  min-height 100vh
  background #fefefe
  box-sizing border-box
  padding 16px 16px 16px 16px
  .search-box
    width 100%
    padding 16px
    box-sizing border-box
    .input
      width 300px
</style>

