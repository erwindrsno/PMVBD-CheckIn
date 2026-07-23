<script setup>
import { ref } from 'vue';
import SchoolList from '@/components/schools/SchoolList.vue';
import EventList from '@/components/events/EventList.vue';
import AttendeeList from '@/components/attendees/AttendeeList.vue';
import Capture from '@/components/attendances/Capture.vue';
import { useRouter } from 'vue-router'

const router = useRouter()

const handleLogout = () => {
  sessionStorage.removeItem('token')

  router.push('/');
}
</script>

<template>
  <v-app-bar rounded>
    <v-app-bar-title>
      <div class="d-flex align-center" style="min-width: 250px;">
        <v-img
          src="../assets/LOGO_PMV.PNG"
          width="50"
          height="50"
          class="mr-3 flex-grow-0"
          cover
        />
        <span>PMVBD, Hadir!</span>
      </div>
    </v-app-bar-title>
    <template v-slot:append>
      <v-btn to="/attendees">Peserta</v-btn>

      <v-menu>
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" append-icon="mdi-chevron-down">
            Kehadiran
          </v-btn>
        </template>

        <v-list>
          <v-list-item to="/attendances/capture" title="Rekam Kehadiran" />
          <v-list-item to="/attendances/list" title="Daftar Kehadiran" />
        </v-list>
      </v-menu>

      <v-btn to="/events">Agenda</v-btn>
      <v-btn to="/schools">Sekolah</v-btn>

      <v-menu>
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" append-icon="mdi-chevron-down">
              User
          </v-btn>
        </template>

        <v-list>
          <!-- Remove the 'to' prop and use @click to trigger a custom logout function -->
          <v-list-item @click="handleLogout" title="Log out" />
        </v-list>
      </v-menu>
    </template>
  </v-app-bar>
</template>
