<template>
  <el-dialog title="条目" :visible.sync="visible">
    <!-- form -->
    <div class="data-form" v-if="data">
      <el-form :model="data" ref="dataForm" label-width="120px" style="width: 90%">
        <el-form-item label="货品编码" prop="goods_id" >
          <el-input v-model="data.goods_id" placeholder="货品编码" disabled  spellcheck="false"></el-input>
        </el-form-item>
        <el-form-item label="名称" prop="name">
          <el-input v-model="data.name" placeholder="请输入名称"  spellcheck="false"></el-input>
        </el-form-item>
        <el-form-item label="名称(英文)" prop="name_e">
          <el-input v-model="data.name_e" placeholder="请输入名称(英文)"  spellcheck="false"></el-input>
        </el-form-item>
        <el-form-item label="地址" prop="addr">
          <el-input v-model="data.addr" placeholder="请输入地址"  spellcheck="false"></el-input>
        </el-form-item>
        <el-form-item label="地址(英文)" prop="addr_e">
          <el-input v-model="data.addr_e" placeholder="请输入地址(英文)" spellcheck="false"></el-input>
        </el-form-item>
      </el-form>
    </div>
    <!-- form:end -->
    <div slot="footer" class="dialog-footer">
      <el-button @click="cancel">取 消</el-button>
      <el-button type="primary" @click="confirm">确 定</el-button>
    </div>
  </el-dialog>
</template>

<script>
import _ from "lodash";

export default {
  data() {
    return {
      visible: false,
      data: null
    };
  },
  methods: {
    gotData(data) {
      this.data = _.clone(data);
    },
    open(data) {
      data && this.gotData(data)
      this.visible = true
    },
    clearData() {
      this.data = null;
    },
    cancel() {
      this.clearData();
      this.visible = false;
    },
    confirm() {
      let cb = action => {
        action = action || Promise.resolve();
        action.then(res => {
          this.cancel();
        });
      };
      this.$emit("on-update-item", this.data, cb);
    }
  }
};
</script>

<style>
</style>
