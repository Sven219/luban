<template>
  <div style="background-color: #FFFFFF">
  <a-tabs v-model:activeKey="data.workload" @change="callback">

    <a-tab-pane key="1" tab="无状态">
      <Deployment></Deployment>
    </a-tab-pane>

    <a-tab-pane key="2" tab="有状态" force-render>
      <StatefulSet></StatefulSet>
    </a-tab-pane>

    <a-tab-pane key="3" tab="守护进程集">
      <DaemonSet></DaemonSet>
    </a-tab-pane>

    <a-tab-pane key="4" tab="任务">
      <Job></Job>
    </a-tab-pane>

    <a-tab-pane key="5" tab="定时任务">
      <CronJob></CronJob>
    </a-tab-pane>

    <a-tab-pane key="6" tab="容器组">
      <Pods></Pods>
    </a-tab-pane>

  </a-tabs>
  </div>
</template>

<script>
import Deployment from "./Deployment";
import Pods from "./Pods";
import {onMounted, reactive} from "vue";
import StatefulSet from "./StatefulSet";
import DaemonSet from "./DaemonSet";
import Job from "./Job";
import CronJob from "./CronJob";
export default {
  name: "WorkLoad",
  setup() {

    const callback = val => {
      localStorage.setItem("workload", val)
    };
    const data = reactive({
          workload: 1,
    })

    const getWorkloadTable = () => {
      data.workload = localStorage.getItem("workload");
      if (data.workload === "" || data.workload === undefined) {
        data.workload = 1
      }
    }

    onMounted(getWorkloadTable)

    return {
      callback,
      data,
    }
  },

  components: {
    CronJob,
    Job,
    DaemonSet,
    Deployment,
    StatefulSet,
    Pods,
  }
}
</script>

<style scoped>

</style>