<script setup>
import { ref } from 'vue';
import SchoolList from './components/schools/SchoolList.vue';
import EventList from './components/events/EventList.vue';
import AttendeeList from './components/attendees/AttendeeList.vue';
import Capture from './components/attendances/Capture.vue';

const currentView = ref('attendee');
</script>

<template>
  <v-app>
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
          <span>PMVBD, Check In!</span>
        </div>
      </v-app-bar-title>
      <template v-slot:append>
        <v-btn to="/attendees">Attendee</v-btn>

        <v-menu>
          <template v-slot:activator="{ props }">
            <v-btn v-bind="props" append-icon="mdi-chevron-down">
              Attendance
            </v-btn>
          </template>

          <v-list>
            <v-list-item to="/attendances/capture" title="Capture Scan" />
            <v-list-item to="/attendances/list" title="Attendance List" />
          </v-list>
        </v-menu>

        <v-btn to="/events">Event</v-btn>
        <v-btn to="/schools">School</v-btn>
      </template>
    </v-app-bar>

    <v-main>
      <v-container fluid>
        <router-view />
      </v-container>
    </v-main>
  </v-app>
</template>
