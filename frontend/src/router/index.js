import { createRouter, createWebHistory } from 'vue-router';
import SchoolList from '@/components/schools/SchoolList.vue';
import EventList from '@/components/events/EventList.vue';
import AttendeeList from '@/components/attendees/AttendeeList.vue';
import Capture from '@/components/attendances/Capture.vue';
import AttendanceList from '@/components/attendances/AttendanceList.vue';
import Login from '@/components/login/Login.vue';
import ProtectedLayout from '@/components/layout/ProtectedLayout.vue';


const routes = [
  {
    path: '/',
    component: Login,
    meta: { title: 'Login | PMVBD' }
  },

  {
    path: '/',
    component: ProtectedLayout,
    meta: { requiresAuth: true },
    children: [
      {
        path: 'attendees',
        component: AttendeeList,
        meta: { title: 'Attendees | PMVBD' }
      },
      {
        path: 'attendances/capture',
        component: Capture,
        meta: { title: 'Capture Scan | PMVBD' }
      },
      {
        path: 'attendances/list',
        component: AttendanceList,
        meta: { title: 'Attendance List | PMVBD' }
      },
      {
        path: 'events',
        component: EventList,
        meta: { title: 'Events | PMVBD' }
      },
      {
        path: 'schools',
        component: SchoolList,
        meta: { title: 'Schools | PMVBD' }
      },
      // Optional: Root redirect to attendees when visiting "/"
      {
        path: '',
        redirect: '/attendees'
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  document.title = to.meta.title || 'PMVBD Check In';

  // 2. Check for the JWT token in local storage
  const token = sessionStorage.getItem('token');

  // 3. Apply the routing rules
  if (to.meta.requiresAuth && !token) {
    // If the page needs auth and there's no token -> send to login
    next('/');
  } else if (to.path === '/' && token) {
    // If user is already logged in and visits the login page -> send to attendees
    next('/attendees');
  } else {
    // Otherwise, proceed normally
    next();
  }
});

export default router;
