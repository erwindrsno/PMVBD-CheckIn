<script setup>
  import { ref, onMounted, computed } from 'vue';
  import ViewDetailDialog from '@/components/attendees/ViewDetailDialog.vue';

  // --- State ---
  const search = ref('');
  const dialog = ref(false);
  const detailDialog = ref(false); // Add this
  const loading = ref(false);
  const attendees = ref([]);
  const selectedAttendee = ref(null);

  // Dropdown Options
  const schools = ref([]);
  const grades = ref([]);
  const subgrades = ref([]);
  const filteredSubgrades = computed(() => {
    // 1. Find the selected grade object
    const selectedGrade = grades.value.find(g => g.id === newAttendee.value.grade_id);

    // 2. If no grade selected or not found, return empty
    if (!selectedGrade) return [];

    // 3. Filter subgrades where subgrade.level_id matches the grade's level_id
    return subgrades.value.filter(s => s.level_id === selectedGrade.level_id);
  });
  const genderOptions = ['Laki-laki', 'Perempuan'];

  const headers = [
    { title: '#', key: 'idx', width: '50px' },
    { title: 'Nama', key: 'name', width: '350px' },
    { title: 'Sekolah', key: 'school'},
    { title: 'Kelas', key: 'grade'},
    { title: 'Telp', key: 'telp' },
    { title: 'Jumlah Kehadiran', key: 'attendance_count', align: 'center'},
    { title: 'Aksi', key: 'action', width:'200px' ,align: 'center' }
  ];

  const viewDetails = (item) => {
    selectedAttendee.value = item;
    detailDialog.value = true;
  };

  const fetchInitialData = async () => {
    loading.value = true;
    try {
      // Fetch both simultaneously
      const token = sessionStorage.getItem('token');
      const [attendeesRes, schoolsRes, gradesRes, subgradesRes] = await Promise.all([
        fetch('http://localhost:8080/api/v1/attendees', {
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
          }
        }),
        fetch('http://localhost:8080/api/v1/schools', {
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
          }
        }), // Replace with your actual schools endpoint
        fetch('http://localhost:8080/api/v1/grades', {
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
          }
        }), // Replace with your actual schools endpoint
        fetch('http://localhost:8080/api/v1/subgrades', {
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
          }
        }) // Replace with your actual schools endpoint
      ]);

      const attendeesData = await attendeesRes.json();
      const schoolsData = await schoolsRes.json();
      const gradesData = await gradesRes.json();
      const subgradesData = await subgradesRes.json();

      attendees.value = attendeesData.data.attendees || [];
      schools.value = schoolsData.data.schools || []; // Assuming your API returns { data: { schools: [...] } }
      grades.value = gradesData.data.grades || [];
      subgrades.value = subgradesData.data.subgrades || [];
    } catch (error) {
      console.error('Error fetching data:', error);
    } finally {
      loading.value = false;
    }
  };

  onMounted(() => {
    fetchInitialData();
  });
</script>

<template>
  <div class="d-flex justify-space-between align-center mb-4">
    <h1>Score Board</h1>

    <div class="d-flex align-center ga-3" style="min-width: 500px;">
      <!-- 2. Search Text Field -->
      <v-text-field
        v-model="search"
        prepend-inner-icon="mdi-magnify"
        label="Cari peserta..."
        variant="outlined"
        density="compact"
        hide-details
        clearable
        single-line
        style="width: 100%;"
      ></v-text-field>

    </div>
  </div>

  <v-data-table
    :headers="headers"
    :items="attendees"
    :loading="loading"
    :search="search"
    :sort-by="[{ key: 'attendance_count', order: 'desc' }]"
    class="elevation-1"
  >
    <!-- Index -->
    <template v-slot:item.idx="{ index }">
      {{ index + 1 }}
    </template>

    <template v-slot:item.grade="{ item }">
      {{ item.grade }} - {{ item.subgrade }}
    </template>

    <!-- Action Column -->
    <template v-slot:header.action>
      <div class="text-center">Aksi</div>
    </template>
    <template v-slot:item.action="{ item }">
      <!-- Call viewDetails(item) instead of editEvent -->
      <v-btn icon variant="text" size="small" color="info" @click="viewDetails(item)">
        <v-icon>mdi-eye</v-icon>
      </v-btn>
    </template>
  </v-data-table>

  <!-- View Detail Dialog -->
  <ViewDetailDialog v-model="detailDialog" :selected-attendee="selectedAttendee" />
</template>
