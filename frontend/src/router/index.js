import { createRouter, createWebHistory } from 'vue-router';
import SchoolList from '../components/schools/SchoolList.vue';
import EventList from '../components/events/EventList.vue';
import AttendeeList from '../components/attendees/AttendeeList.vue';
import Capture from '../components/attendances/Capture.vue';
import AttendanceList from '../components/attendances/AttendanceList.vue';

const routes = [
  {
    path: '/attendees',
    component: AttendeeList,
    meta: { title: 'Attendees | PMVBD' }
  },
  {
    path: '/attendances/capture',
    component: Capture,
    meta: { title: 'Capture Scan | PMVBD' }
  },
  {
    path: '/attendances/list',
    component: AttendanceList,
    meta: { title: 'Attendance List | PMVBD' }
  },
  {
    path: '/events',
    component: EventList,
    meta: { title: 'Events | PMVBD' }
  },
  {
    path: '/schools',
    component: SchoolList,
    meta: { title: 'Schools | PMVBD' }
  },
  {
    path: '/',
    redirect: '/attendees'
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  document.title = to.meta.title || 'PMVBD Check In';
  next();
});

export default router;
