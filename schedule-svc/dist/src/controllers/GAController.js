"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.GAScheduleController = void 0;
const lodash_1 = __importDefault(require("lodash"));
// import { RoomDefModel } from "../models/roomDefModel";
// import { CouncilModel } from "../models/councilModel";
// import { StudentDefModel } from "../models/studentDefModel";
const moment_1 = __importDefault(require("moment"));
const uuid_1 = require("uuid");
const data_1 = require("../data");
const scheduledModel_1 = require("../models/scheduledModel");
exports.GAScheduleController = {
    // async scheduleGAThesisDefense(req: Request, res: Response) {
    //   try {
    //     const scheduleParams: IParamSchedule = req.body
    //     // Lấy dữ liệu từ database
    //     async function getAllRoomDef() {
    //       try {
    //         const response = await RoomDefModel.getAllRoomDef();
    //         const roomList = response;
    //         return roomList;
    //       } catch (error) {
    //         console.error(error);
    //         return [];
    //       }
    //     }
    //     async function getAllCouncilDef() {
    //       try {
    //         const response = await CouncilModel.getAllCouncil();
    //         const councilList = response;
    //         return councilList;
    //       } catch (error) {
    //         console.error(error);
    //         return [];
    //       }
    //     }
    //     async function getAllStudentDef() {
    //       try {
    //         const response = await StudentDefModel.getAllStudentDef();
    //         const studentList = response;
    //         return studentList;
    //       } catch (error) {
    //         console.error(error);
    //         return [];
    //       }
    //     }
    //     const roomLists = await getAllRoomDef();
    //     const councilLists = await getAllCouncilDef();
    //     const studentLists = await getAllStudentDef();
    //     // Function handle number of rooms
    //     const handleNumberRoomDef = (wks: number, studef: number) => {
    //       if (wks <= 0) {
    //         return 0;
    //       } else {
    //         const convertDate = wks * 6;
    //         const dateDefense = studef / convertDate;
    //         return Math.ceil(dateDefense / 12);
    //       }
    //     };
    //     // Định nghĩa tham số cho thuật toán di truyền
    //     const populationSize: number = 10; // Kích thước quần thể
    //     const maxGenerations: number = 5; // Số thế hệ tối đa
    //     const mutationRate: number = 0.01; // Tỷ lệ đột biến
    //     const numLecturers: number = councilLists.length; // Số lượng giảng viên
    //     const numRooms: number = roomLists.length; // Số lượng phòng
    //     const numStudent: number = studentLists.length; // Số lượng sinh viên
    //     // Khởi tạo quần thể ban đầu
    //     let initPopulation: IThesisDef[] = initializePopulation(populationSize);
    //     // Hàm khởi tạo quần thể ban đầu
    //     function initializePopulation(size: number): IThesisDef[] {
    //       const population: IThesisDef[] = [];
    //       for (let i = 0; i < size; i++) {
    //         const schedule = generateCouncilRoom();
    //         population.push({ thesis: schedule, fitness: 1 });
    //       }
    //       return population;
    //     }
    //     function getNumberOfDays(startTime: string, endTime: string): number {
    //       const startDate = moment(startTime);
    //       const endDate = moment(endTime);
    //       const duration = moment.duration(endDate.diff(startDate));
    //       const numberOfDays = duration.asDays();
    //       return numberOfDays;
    //     }
    //     function checkQuantityStudentValid(
    //       lecDef: IAuthObject[],
    //       studefArr: IStudentDef[],
    //       reportStudef: string[]
    //     ): boolean {
    //       let stateDef = false;
    //       for (let i = 0; i < numLecturers; i++) {
    //         const studentFollowLecturer = studefArr.filter((student) => {
    //           return (
    //             lecDef[i].id === student.instructor.id &&
    //             !reportStudef.includes(student.id)
    //           );
    //         });
    //         if (studentFollowLecturer.length >= 10) {
    //           stateDef = true;
    //         }
    //       }
    //       return stateDef;
    //     }
    //     function getLecturerListMore10(
    //       lecDef: IAuthObject[],
    //       studefArr: IStudentDef[],
    //       reportStudef: string[]
    //     ): IAuthObject[] {
    //       let lecturerHasAvailbleStudef: IAuthObject[] = [];
    //       for (let i = 0; i < numLecturers; i++) {
    //         const studentFollowLecturer = studefArr.filter((student) => {
    //           return (
    //             lecDef[i].id === student.instructor.id &&
    //             !reportStudef.includes(student.id)
    //           );
    //         });
    //         if (studentFollowLecturer.length >= 10) {
    //           lecturerHasAvailbleStudef.push(lecDef[i]);
    //         }
    //       }
    //       return lecturerHasAvailbleStudef;
    //     }
    //     function getLecturerListlend(
    //       lecDef: IAuthObject[],
    //       studefArr: IStudentDef[],
    //       reportStudef: string[]
    //     ): IAuthObject[] {
    //       let lecturerHasAvailbleStudef: IAuthObject[] = [];
    //       for (let i = 0; i < numLecturers; i++) {
    //         const studentFollowLecturer = studefArr.filter((student) => {
    //           return (
    //             lecDef[i].id === student.instructor.id &&
    //             !reportStudef.includes(student.id)
    //           );
    //         });
    //         if (studentFollowLecturer.length > 0) {
    //           lecturerHasAvailbleStudef.push(lecDef[i]);
    //         }
    //       }
    //       return lecturerHasAvailbleStudef;
    //     }
    //     // Hàm đưa ds giảng viên (hội đông) gổm 3 người vào phòng
    //     function generateCouncilRoom(): ICouncilDef[] {
    //       let currentDate = new Date(scheduleParams.startDate);
    //       let councilSchedules: ICouncilDef[] = [];
    //       let reportedStudents: string[] = []; // Danh sách SV đã báo cáo
    //       while (true) {
    //         const shuffledLecturers: IAuthObject[] = shuffleArray(councilLists);
    //         const shuffledStudents: IStudentDef[] = shuffleArray(studentLists);
    //         let reportedLecturers: IAuthObject[] = []; // Danh sách GV đã báo cáo
    //         let roomIndex = 0; // Biến đánh dấu chỉ số của phòng
    //         // Lặp qua danh sách lấy ra báo cáo theo lớp 12 sv / 15sv lớp
    //         for (
    //           let i = 0;
    //           i < numLecturers;
    //           i += handleNumberRoomDef(scheduleParams.quantityWeek, numStudent)
    //         ) {
    //           let oldCouncil: IAuthObject[] = [];
    //           let listOldLecturer: IAuthObject[] = [];
    //           for (
    //             let j = 0;
    //             j < handleNumberRoomDef(scheduleParams.quantityWeek, numStudent);
    //             j++
    //           ) {
    //             let councilEstab: IAuthObject[] = [];
    //             const availableLecturers = shuffledLecturers.filter((lecturer) => !reportedLecturers.includes(lecturer));
    //             if (availableLecturers.length > 0) {
    //               if (currentDate.getDay() === 0) {
    //                 currentDate.setDate(currentDate.getDate() + 1);
    //               }
    //               const roomSchedules = generateScheduleRoom(currentDate);
    //               const councilLecturer = availableLecturers[0];
    //               reportedLecturers.push(councilLecturer);
    //               listOldLecturer.push(councilLecturer);
    //               const getMoreLecturers = shuffleArray(shuffledLecturers)
    //                 .filter(
    //                   (lecturer) =>
    //                     !oldCouncil.includes(lecturer) &&
    //                     !listOldLecturer.includes(lecturer)
    //                 )
    //                 .slice(0, 2);
    //               oldCouncil.push(...getMoreLecturers, councilLecturer);
    //               councilEstab.push(...getMoreLecturers, councilLecturer);
    //               const studentFollowLecturer = shuffledStudents.filter(
    //                 (student) => {
    //                   return (
    //                     councilLecturer.id === student.instructor.id &&
    //                     !reportedStudents.includes(student.id)
    //                   );
    //                 }
    //               );
    //               if (studentFollowLecturer.length >= 10) {
    //                 councilSchedules.push({
    //                   id: uuidv4(),
    //                   council: councilEstab,
    //                   schedule: {
    //                     room: roomSchedules[j].room,
    //                     timeSlots: assignStudentsToTimeSlots(
    //                       roomSchedules[j].timeSlots,
    //                       studentFollowLecturer
    //                     ),
    //                   },
    //                 });
    //                 const reportedStudentIds = studentFollowLecturer
    //                   .slice(0, 12)
    //                   .map((student) => student.id);
    //                 reportedStudents.push(...reportedStudentIds);
    //               }
    //             } else {
    //               break;
    //             }
    //             roomIndex = (roomIndex + 1) % 2;
    //             if (roomIndex === 0) {
    //               currentDate.setDate(currentDate.getDate() + 1);
    //             }
    //           }
    //         }
    //         currentDate.setDate(currentDate.getDate() + 1);
    //         if (
    //           !checkQuantityStudentValid(
    //             shuffledLecturers,
    //             shuffledStudents,
    //             reportedStudents
    //           )
    //         ) {
    //           break;
    //         }
    //       }
    //       // Tiếp tục lặp qua danh sách thừa để lặp lịch 1 lần nữa hội đồng sẽ là 3 gv đều có sinh viên thừa
    //       while (true) {
    //         let roomIndex = 0;
    //         let reportedLecturerSecond: IAuthObject[] = [];
    //         const shuffledLecturerSeconds: IAuthObject[] =
    //           shuffleArray(councilLists);
    //         const shuffledStudentSeconds: IStudentDef[] =
    //           shuffleArray(studentLists);
    //         const availableStudents = shuffledStudentSeconds.filter(
    //           (student) => !reportedStudents.includes(student.id)
    //         );
    //         for (let k = 0; k <= numLecturers; k += 6) {
    //           for (
    //             let i = 0;
    //             i < handleNumberRoomDef(scheduleParams.quantityWeek, numStudent);
    //             i++
    //           ) {
    //             const councilLecturers = getLecturerListlend(
    //               shuffledLecturerSeconds,
    //               shuffledStudentSeconds,
    //               reportedStudents
    //             )
    //               .filter((lec) => !reportedLecturerSecond.includes(lec))
    //               .slice(0, 3);
    //             if (councilLecturers.length != 3) {
    //               let newcouncilLecturers = reportedLecturerSecond.slice(0, 3 - councilLecturers.length);
    //               councilLecturers.push(...newcouncilLecturers);
    //             }
    //             reportedLecturerSecond.push(...councilLecturers);
    //             if (currentDate.getDay() === 0) {
    //               currentDate.setDate(currentDate.getDate() + 1);
    //             }
    //             const roomSchedules = generateScheduleRoomSecond(currentDate);
    //             const studentFollowLecturer = availableStudents.filter(
    //               (student) => {
    //                 return (
    //                   councilLecturers.some(
    //                     (lecturer) => lecturer.id === student.instructor.id
    //                   ) && !reportedStudents.includes(student.id)
    //                 );
    //               }
    //             );
    //             if (studentFollowLecturer.length > 0) {
    //               councilSchedules.push({
    //                 id: uuidv4(),
    //                 council: councilLecturers,
    //                 schedule: {
    //                   room: roomSchedules[i].room,
    //                   timeSlots: assignStudentsToTimeSlots(
    //                     roomSchedules[i].timeSlots,
    //                     studentFollowLecturer
    //                   ),
    //                 },
    //               });
    //               const reportedStudentIds = studentFollowLecturer
    //                 .slice(0, 10)
    //                 .map((student) => student.id);
    //               reportedStudents.push(...reportedStudentIds);
    //             }
    //             roomIndex = (roomIndex + 1) % 2;
    //             if (roomIndex === 0) {
    //               currentDate.setDate(currentDate.getDate() + 1);
    //             }
    //           }
    //         }
    //         if (
    //           !getLecturerListlend(
    //             shuffledLecturerSeconds,
    //             shuffledStudentSeconds,
    //             reportedStudents
    //           ).length
    //         ) {
    //           break;
    //         }
    //       }
    //       return councilSchedules;
    //     }
    //     // Hàm add student vào timeslot tương ứng
    //     function assignStudentsToTimeSlots(
    //       timeSlots: ITimeSlotItem[],
    //       students: IStudentDef[]
    //     ) {
    //       const assignStudentArr: ITimeSlotForStudent[] = [];
    //       const emptyStudent: IStudentDef = INITIATE_STUDENT_DEF;
    //       for (let i = 0; i < timeSlots.length; i++) {
    //         const student = i < students.length ? students[i] : emptyStudent;
    //         assignStudentArr.push({
    //           timeSlot: timeSlots[i],
    //           student: student,
    //         });
    //       }
    //       return assignStudentArr;
    //     }
    //     // Hàm trộn danh sách giảng viên
    //     function shuffleArray(array: any[]): any[] {
    //       const shuffledArray = [...array];
    //       for (let i = shuffledArray.length - 1; i > 0; i--) {
    //         const j = Math.floor(Math.random() * (i + 1));
    //         [shuffledArray[i], shuffledArray[j]] = [
    //           shuffledArray[j],
    //           shuffledArray[i],
    //         ];
    //       }
    //       return shuffledArray;
    //     }
    //     // Hàm tạo danh sách phòng và lịch trình cho mỗi phòng cố định
    //     function generateScheduleRoom(
    //       currentDate: Date
    //     ): IScheduleDefForStudent[] {
    //       const roomSchedules: IScheduleDefForStudent[] = [];
    //       roomLists
    //         .slice(
    //           0,
    //           handleNumberRoomDef(scheduleParams.quantityWeek, numStudent)
    //         )
    //         .map((room:any) => {
    //           const timeSlots = generateTimeSlots(currentDate);
    //           roomSchedules.push({
    //             room: room,
    //             timeSlots: timeSlots,
    //           });
    //         });
    //       return roomSchedules;
    //     }
    //     // Hàm tạo danh sách thời gian báo cáo cụ thể (12 people/day)
    //     function generateTimeSlots(currentDate: Date): ITimeSlotItem[] {
    //       const timeSlots: ITimeSlotItem[] = [];
    //       const minutesPerSlot = 40; //Input
    //       const timeSlotCountMor = 7; // Input
    //       const timeSlotCountAf = 5; // Input
    //       const formattedDate = currentDate.toLocaleDateString("en-US", {
    //         weekday: "long",
    //         month: "numeric",
    //         day: "numeric",
    //         year: "numeric",
    //       });
    //       for (let am = 0; am < timeSlotCountMor; am++) {
    //         const startTime = new Date(currentDate);
    //         startTime.setHours(7, am * minutesPerSlot);
    //         for (let k = 0; k < minutesPerSlot; k += minutesPerSlot) {
    //           const timeSlot = {
    //             id: uuidv4(),
    //             date: formattedDate,
    //             time: startTime.toLocaleTimeString("en-US", {
    //               hour: "numeric",
    //               minute: "numeric",
    //             }),
    //             shift: "Morning",
    //           };
    //           timeSlots.push(timeSlot);
    //           startTime.setMinutes(startTime.getMinutes() + minutesPerSlot);
    //         }
    //       }
    //       for (let pm = 0; pm < timeSlotCountAf; pm++) {
    //         const startTime = new Date(currentDate);
    //         startTime.setHours(13, 30 + pm * minutesPerSlot);
    //         for (let k = 0; k < minutesPerSlot; k += minutesPerSlot) {
    //           const timeSlot = {
    //             id: uuidv4(),
    //             date: formattedDate,
    //             time: startTime.toLocaleTimeString("en-US", {
    //               hour: "numeric",
    //               minute: "numeric",
    //             }),
    //             shift: "Afternoon",
    //           };
    //           timeSlots.push(timeSlot);
    //           startTime.setMinutes(startTime.getMinutes() + minutesPerSlot);
    //         }
    //       }
    //       return timeSlots;
    //     }
    //     // Hàm tạo danh sách phòng và lịch trình cho mỗi phòng cố định
    //     function generateScheduleRoomSecond(
    //       currentDate: Date
    //     ): IScheduleDefForStudent[] {
    //       const roomSchedules: IScheduleDefForStudent[] = [];
    //       roomLists
    //         .slice(
    //           0,
    //           handleNumberRoomDef(scheduleParams.quantityWeek, numStudent)
    //         )
    //         .map((room:any) => {
    //           const timeSlots = generateTimeSlotSeconds(currentDate);
    //           roomSchedules.push({
    //             room: room,
    //             timeSlots: timeSlots,
    //           });
    //         });
    //       return roomSchedules;
    //     }
    //     // Hàm tạo danh sách thời gian báo cáo cụ thể (10 people/day)
    //     function generateTimeSlotSeconds(currentDate: Date): ITimeSlotItem[] {
    //       const timeSlots: ITimeSlotItem[] = [];
    //       const minutesPerSlot = 40; //Input
    //       const timeSlotCountMor = 5; // Input
    //       const timeSlotCountAf = 5; // Input
    //       const formattedDate = currentDate.toLocaleDateString("en-US", {
    //         weekday: "long",
    //         month: "numeric",
    //         day: "numeric",
    //         year: "numeric",
    //       });
    //       for (let am = 0; am < timeSlotCountMor; am++) {
    //         const startTime = new Date(currentDate);
    //         startTime.setHours(7, am * minutesPerSlot);
    //         for (let k = 0; k < minutesPerSlot; k += minutesPerSlot) {
    //           const timeSlot = {
    //             id: uuidv4(),
    //             date: formattedDate,
    //             time: startTime.toLocaleTimeString("en-US", {
    //               hour: "numeric",
    //               minute: "numeric",
    //             }),
    //             shift: "Morning",
    //           };
    //           timeSlots.push(timeSlot);
    //           startTime.setMinutes(startTime.getMinutes() + minutesPerSlot);
    //         }
    //       }
    //       for (let pm = 0; pm < timeSlotCountAf; pm++) {
    //         const startTime = new Date(currentDate);
    //         startTime.setHours(13, 30 + pm * minutesPerSlot);
    //         for (let k = 0; k < minutesPerSlot; k += minutesPerSlot) {
    //           const timeSlot = {
    //             id: uuidv4(),
    //             date: formattedDate,
    //             time: startTime.toLocaleTimeString("en-US", {
    //               hour: "numeric",
    //               minute: "numeric",
    //             }),
    //             shift: "Afternoon",
    //           };
    //           timeSlots.push(timeSlot);
    //           startTime.setMinutes(startTime.getMinutes() + minutesPerSlot);
    //         }
    //       }
    //       return timeSlots;
    //     }
    //     // Hàm mục tiêu, tính toán độ chính xác dựa vào lịch trình nếu có xảy ra conflict
    //     function fitness(thesis: IThesisDef, totalStudents: number): number {
    //       let fitnessScore = 0;
    //       function findStudentConflicts(thesis: IThesisDef) {
    //         const conflicts: ITimeSlotForStudent[] = [];
    //         const timeSlots: IStudentDef[] = [];
    //         for (let i = 0; i < thesis.thesis.length; i++) {
    //           const timeSlotsData = thesis.thesis[i].schedule.timeSlots;
    //           for (let j = 0; j < timeSlotsData.length; j++) {
    //             const timeSlot = timeSlotsData[j];
    //             const student = timeSlot.student;
    //             if (timeSlots.includes(student) && student.id !== "") {
    //               conflicts.push({
    //                 student: timeSlot.student,
    //                 timeSlot: timeSlot.timeSlot,
    //               });
    //             } else {
    //               timeSlots.push(student);
    //             }
    //           }
    //         }
    //         return conflicts;
    //       }
    //       const studentConflicts = findStudentConflicts(thesis);
    //       // console.log(studentConflicts.length);
    //       const numStudentConflicts = studentConflicts.length;
    //       const conflictRatio = numStudentConflicts / totalStudents;
    //       fitnessScore += 1 - conflictRatio;
    //       return fitnessScore;
    //     }
    //     // Lấy 70% lịch trình trong quần thể để tiến hành lai ghép và đột biến
    //     function selection(evaluatedPopulation: IThesisDef[]): IThesisDef[] {
    //       const populationSize = evaluatedPopulation.length;
    //       const totalFitness = evaluatedPopulation.reduce(
    //         (sum, individual) => sum + individual.fitness,
    //         0
    //       );
    //       const selectionProbabilities = evaluatedPopulation.map(
    //         (individual) => individual.fitness / totalFitness
    //       );
    //       const selectedPopulation: IThesisDef[] = [];
    //       const selectedPopulationSize = Math.floor(populationSize * 0.7); // Chọn 70% số lượng cá thể
    //       for (let i = 0; i < selectedPopulationSize; i++) {
    //         let selectedIndex = 0;
    //         let randomValue = Math.random();
    //         while (randomValue > 0) {
    //           randomValue -= selectionProbabilities[selectedIndex];
    //           selectedIndex++;
    //         }
    //         selectedIndex--;
    //         selectedPopulation.push(evaluatedPopulation[selectedIndex]);
    //       }
    //       return selectedPopulation;
    //     }
    //     // Qúa trình lai ghép 2 cá thể trong quần thể
    //     function crossover(
    //       parent1?: IThesisDef,
    //       parent2?: IThesisDef
    //     ): IThesisDef[] {
    //       // Copy information from parents
    //       const child1: IThesisDef = JSON.parse(JSON.stringify(parent1));
    //       const child2: IThesisDef = JSON.parse(JSON.stringify(parent2));
    //       // Select two random student positions in the child individuals
    //       const studentIndex1 = Math.floor(
    //         Math.random() * child1.thesis[0].schedule.timeSlots.length
    //       );
    //       const studentIndex2 = Math.floor(
    //         Math.random() * child2.thesis[0].schedule.timeSlots.length
    //       );
    //       // Check if both objects have values
    //       let student1 =
    //         child1.thesis[0].schedule.timeSlots[studentIndex1].student;
    //       let student2 =
    //         child2.thesis[0].schedule.timeSlots[studentIndex2].student;
    //       if (student1 && student2) {
    //         const tempStudent = student1;
    //         student1 = student2;
    //         student2 = tempStudent;
    //       }
    //       // Return the array after swapping
    //       return [child1, child2];
    //     }
    //     // Qúa trình đột biến diễn ra trong quần thể với tỷ lệ 10% tức là 1-50 sv thì sẽ có đột biến
    //     function mutation(individual: IThesisDef, mutationRate: number) {
    //       const timetableLength = individual.thesis.length;
    //       for (let i = 0; i < timetableLength; i++) {
    //         if (Math.random() < mutationRate) {
    //           const thesis = individual.thesis[i];
    //           const timeSlotsLength = thesis.schedule.timeSlots.length;
    //           if (timeSlotsLength > 0) {
    //             const randomIndex = Math.floor(Math.random() * timeSlotsLength);
    //             const startDate = new Date("November 27, 2023");
    //             let timeSlot = thesis.schedule.timeSlots[randomIndex];
    //             const currentDate = new Date(startDate);
    //             const startTime = new Date(currentDate);
    //             currentDate.setDate(startDate.getDate() + i);
    //             const formattedDate = currentDate.toLocaleDateString("en-US", {
    //               weekday: "long",
    //               month: "numeric",
    //               day: "numeric",
    //               year: "numeric",
    //             });
    //             startTime.setHours(7, 0);
    //             let newTimeSlot = {
    //               id: uuidv4(),
    //               date: formattedDate,
    //               time: startTime.toLocaleTimeString("en-US", {
    //                 hour: "numeric",
    //                 minute: "numeric",
    //               }),
    //               shift: "Morning",
    //             };
    //             timeSlot.timeSlot = newTimeSlot;
    //           }
    //         }
    //       }
    //       return individual;
    //     }
    //     // Chạy giải thuật Genetic algorithm
    //     let mutionPopulation: IThesisDef[] = [];
    //     for (let generation = 0; generation < maxGenerations; generation++) {
    //       // Đánh giá độ thích nghi của quần thể
    //       const totalScheduleTimeSlots = initPopulation.reduce(
    //         (total, individual) => {
    //           const individualTimeSlots = individual.thesis.reduce(
    //             (thesisTotal: number, thesis: ICouncilDef) => {
    //               return thesisTotal + thesis.schedule.timeSlots.length;
    //             },
    //             0
    //           );
    //           return total + individualTimeSlots;
    //         },
    //         0
    //       );
    //       const evaluatedPopulation = initPopulation.map(
    //         (individual, index) =>
    //           ({
    //             thesis: individual.thesis,
    //             fitness: fitness(individual, totalScheduleTimeSlots),
    //           } as IThesisDef)
    //       );
    //       // Lựa chọn cá thể cho thế hệ mới
    //       const selectedPopulation = selection(evaluatedPopulation);
    //       // Tạo thế hệ mới bằng cách lai ghép và đột biến cá thể
    //       let newPopulation: IThesisDef[] = [];
    //       while (newPopulation.length < populationSize) {
    //         // Lựa chọn hai cá thể cha mẹ
    //         const parent1 = _.sample(selectedPopulation);
    //         const parent2 = _.sample(selectedPopulation);
    //         // Lai ghép để tạo ra cá thể con
    //         const child = crossover(parent1, parent2);
    //         // Đột biến cá thể con
    //         for (let i = 0; i < child.length; i++) {
    //           const newChild = mutation(child[i], mutationRate);
    //           newPopulation.push(newChild);
    //         }
    //       }
    //       // Cập nhật quần thể mới
    //       mutionPopulation = newPopulation;
    //     }
    //     // Sắp xếp giảm dần theo độ chính xác
    //     mutionPopulation.sort((a, b) => {
    //       if (a.fitness !== b.fitness) {
    //         return b.fitness - a.fitness; // Sắp xếp theo fitness cao nhất
    //       } else {
    //         return a.thesis.length - b.thesis.length; // Nếu fitness giống nhau, sắp xếp theo độ dài bestSchedule tăng dần
    //       }
    //     });
    //     // Lấy ra lịch trình trong quần thể (Với độ chính xác cao nhất và độ dài thấp nhất)
    //     const bestSchedule: ICouncilDef[] = mutionPopulation[0].thesis;
    //     // Trả về client lịch trình tốt nhất
    //     res.status(200).json(bestSchedule);
    //     await ScheduleDefModel.saveScheduleDef({
    //       thesis: bestSchedule,
    //     } as IThesisDef);
    //   } catch (err) {
    //     console.log(err);
    //   }
    // },
    async scheduleGAThesisDefense(req) {
        try {
            const scheduleParams = {
                startDate: req.request.startTime,
                quantityWeek: req.request.quantityWeek,
            };
            const roomLists = req.request.rooms;
            const councilLists = req.request.councils;
            const studentLists = req.request.studentDefs;
            // Function handle number of rooms
            const handleNumberRoomDef = (wks, studef) => {
                if (wks <= 0) {
                    return 0;
                }
                else {
                    const convertDate = wks * 6;
                    const dateDefense = studef / convertDate;
                    return Math.ceil(dateDefense / 12);
                }
            };
            // Định nghĩa tham số cho thuật toán di truyền
            const populationSize = 10; // Kích thước quần thể
            const maxGenerations = 5; // Số thế hệ tối đa
            const mutationRate = 0.01; // Tỷ lệ đột biến
            const numLecturers = councilLists.length; // Số lượng giảng viên
            const numRooms = roomLists.length; // Số lượng phòng
            const numStudent = studentLists.length; // Số lượng sinh viên
            // Khởi tạo quần thể ban đầu
            let initPopulation = initializePopulation(populationSize);
            // Hàm khởi tạo quần thể ban đầu
            function initializePopulation(size) {
                const population = [];
                for (let i = 0; i < size; i++) {
                    const schedule = generateCouncilRoom();
                    population.push({ thesis: schedule, fitness: 1 });
                }
                return population;
            }
            function getNumberOfDays(startTime, endTime) {
                const startDate = (0, moment_1.default)(startTime);
                const endDate = (0, moment_1.default)(endTime);
                const duration = moment_1.default.duration(endDate.diff(startDate));
                const numberOfDays = duration.asDays();
                return numberOfDays;
            }
            function checkQuantityStudentValid(lecDef, studefArr, reportStudef) {
                let stateDef = false;
                for (let i = 0; i < numLecturers; i++) {
                    const studentFollowLecturer = studefArr.filter((student) => {
                        return (lecDef[i].id === student.instructor.id &&
                            !reportStudef.includes(student.id));
                    });
                    if (studentFollowLecturer.length >= 10) {
                        stateDef = true;
                    }
                }
                return stateDef;
            }
            function getLecturerListMore10(lecDef, studefArr, reportStudef) {
                let lecturerHasAvailbleStudef = [];
                for (let i = 0; i < numLecturers; i++) {
                    const studentFollowLecturer = studefArr.filter((student) => {
                        return (lecDef[i].id === student.instructor.id &&
                            !reportStudef.includes(student.id));
                    });
                    if (studentFollowLecturer.length >= 10) {
                        lecturerHasAvailbleStudef.push(lecDef[i]);
                    }
                }
                return lecturerHasAvailbleStudef;
            }
            function getLecturerListlend(lecDef, studefArr, reportStudef) {
                let lecturerHasAvailbleStudef = [];
                for (let i = 0; i < numLecturers; i++) {
                    const studentFollowLecturer = studefArr.filter((student) => {
                        return (lecDef[i].id === student.instructor.id &&
                            !reportStudef.includes(student.id));
                    });
                    if (studentFollowLecturer.length > 0) {
                        lecturerHasAvailbleStudef.push(lecDef[i]);
                    }
                }
                return lecturerHasAvailbleStudef;
            }
            // Hàm đưa ds giảng viên (hội đông) gổm 3 người vào phòng
            function generateCouncilRoom() {
                let currentDate = new Date(scheduleParams.startDate);
                let councilSchedules = [];
                let reportedStudents = []; // Danh sách SV đã báo cáo
                while (true) {
                    const shuffledLecturers = shuffleArray(councilLists);
                    const shuffledStudents = shuffleArray(studentLists);
                    let reportedLecturers = []; // Danh sách GV đã báo cáo
                    let roomIndex = 0; // Biến đánh dấu chỉ số của phòng
                    // Lặp qua danh sách lấy ra báo cáo theo lớp 12 sv / 15sv lớp
                    for (let i = 0; i < numLecturers; i += handleNumberRoomDef(scheduleParams.quantityWeek, numStudent)) {
                        let oldCouncil = [];
                        let listOldLecturer = [];
                        for (let j = 0; j < handleNumberRoomDef(scheduleParams.quantityWeek, numStudent); j++) {
                            let councilEstab = [];
                            const availableLecturers = shuffledLecturers.filter((lecturer) => !reportedLecturers.includes(lecturer));
                            if (availableLecturers.length > 0) {
                                if (currentDate.getDay() === 0) {
                                    currentDate.setDate(currentDate.getDate() + 1);
                                }
                                const roomSchedules = generateScheduleRoom(currentDate);
                                const councilLecturer = availableLecturers[0];
                                reportedLecturers.push(councilLecturer);
                                listOldLecturer.push(councilLecturer);
                                const getMoreLecturers = shuffleArray(shuffledLecturers)
                                    .filter((lecturer) => !oldCouncil.includes(lecturer) &&
                                    !listOldLecturer.includes(lecturer))
                                    .slice(0, 2);
                                oldCouncil.push(...getMoreLecturers, councilLecturer);
                                councilEstab.push(...getMoreLecturers, councilLecturer);
                                const studentFollowLecturer = shuffledStudents.filter((student) => {
                                    return (councilLecturer.id === student.instructor.id &&
                                        !reportedStudents.includes(student.id));
                                });
                                if (studentFollowLecturer.length >= 10) {
                                    councilSchedules.push({
                                        id: (0, uuid_1.v4)(),
                                        council: councilEstab,
                                        schedule: {
                                            room: roomSchedules[j].room,
                                            timeSlots: assignStudentsToTimeSlots(roomSchedules[j].timeSlots, studentFollowLecturer),
                                        },
                                    });
                                    const reportedStudentIds = studentFollowLecturer
                                        .slice(0, 12)
                                        .map((student) => student.id);
                                    reportedStudents.push(...reportedStudentIds);
                                }
                            }
                            else {
                                break;
                            }
                            roomIndex = (roomIndex + 1) % 2;
                            if (roomIndex === 0) {
                                currentDate.setDate(currentDate.getDate() + 1);
                            }
                        }
                    }
                    currentDate.setDate(currentDate.getDate() + 1);
                    if (!checkQuantityStudentValid(shuffledLecturers, shuffledStudents, reportedStudents)) {
                        break;
                    }
                }
                // Tiếp tục lặp qua danh sách thừa để lặp lịch 1 lần nữa hội đồng sẽ là 3 gv đều có sinh viên thừa
                while (true) {
                    let roomIndex = 0;
                    let reportedLecturerSecond = [];
                    const shuffledLecturerSeconds = shuffleArray(councilLists);
                    const shuffledStudentSeconds = shuffleArray(studentLists);
                    const availableStudents = shuffledStudentSeconds.filter((student) => !reportedStudents.includes(student.id));
                    for (let k = 0; k <= numLecturers; k += 6) {
                        for (let i = 0; i < handleNumberRoomDef(scheduleParams.quantityWeek, numStudent); i++) {
                            const councilLecturers = getLecturerListlend(shuffledLecturerSeconds, shuffledStudentSeconds, reportedStudents)
                                .filter((lec) => !reportedLecturerSecond.includes(lec))
                                .slice(0, 3);
                            if (councilLecturers.length != 3) {
                                let newcouncilLecturers = reportedLecturerSecond.slice(0, 3 - councilLecturers.length);
                                councilLecturers.push(...newcouncilLecturers);
                            }
                            reportedLecturerSecond.push(...councilLecturers);
                            if (currentDate.getDay() === 0) {
                                currentDate.setDate(currentDate.getDate() + 1);
                            }
                            const roomSchedules = generateScheduleRoomSecond(currentDate);
                            const studentFollowLecturer = availableStudents.filter((student) => {
                                return (councilLecturers.some((lecturer) => lecturer.id === student.instructor.id) && !reportedStudents.includes(student.id));
                            });
                            if (studentFollowLecturer.length > 0) {
                                councilSchedules.push({
                                    id: (0, uuid_1.v4)(),
                                    council: councilLecturers,
                                    schedule: {
                                        room: roomSchedules[i].room,
                                        timeSlots: assignStudentsToTimeSlots(roomSchedules[i].timeSlots, studentFollowLecturer),
                                    },
                                });
                                const reportedStudentIds = studentFollowLecturer
                                    .slice(0, 10)
                                    .map((student) => student.id);
                                reportedStudents.push(...reportedStudentIds);
                            }
                            roomIndex = (roomIndex + 1) % 2;
                            if (roomIndex === 0) {
                                currentDate.setDate(currentDate.getDate() + 1);
                            }
                        }
                    }
                    if (!getLecturerListlend(shuffledLecturerSeconds, shuffledStudentSeconds, reportedStudents).length) {
                        break;
                    }
                }
                return councilSchedules;
            }
            // Hàm add student vào timeslot tương ứng
            function assignStudentsToTimeSlots(timeSlots, students) {
                const assignStudentArr = [];
                const emptyStudent = data_1.INITIATE_STUDENT_DEF;
                for (let i = 0; i < timeSlots.length; i++) {
                    const student = i < students.length ? students[i] : emptyStudent;
                    assignStudentArr.push({
                        timeSlot: timeSlots[i],
                        student: student,
                    });
                }
                return assignStudentArr;
            }
            // Hàm trộn danh sách giảng viên
            function shuffleArray(array) {
                const shuffledArray = [...array];
                for (let i = shuffledArray.length - 1; i > 0; i--) {
                    const j = Math.floor(Math.random() * (i + 1));
                    [shuffledArray[i], shuffledArray[j]] = [
                        shuffledArray[j],
                        shuffledArray[i],
                    ];
                }
                return shuffledArray;
            }
            // Hàm tạo danh sách phòng và lịch trình cho mỗi phòng cố định
            function generateScheduleRoom(currentDate) {
                const roomSchedules = [];
                roomLists
                    .slice(0, handleNumberRoomDef(scheduleParams.quantityWeek, numStudent))
                    .map((room) => {
                    const timeSlots = generateTimeSlots(currentDate);
                    roomSchedules.push({
                        room: room,
                        timeSlots: timeSlots,
                    });
                });
                return roomSchedules;
            }
            // Hàm tạo danh sách thời gian báo cáo cụ thể (12 people/day)
            function generateTimeSlots(currentDate) {
                const timeSlots = [];
                const minutesPerSlot = 40; //Input
                const timeSlotCountMor = 7; // Input
                const timeSlotCountAf = 5; // Input
                const formattedDate = currentDate.toLocaleDateString("en-US", {
                    weekday: "long",
                    month: "numeric",
                    day: "numeric",
                    year: "numeric",
                });
                for (let am = 0; am < timeSlotCountMor; am++) {
                    const startTime = new Date(currentDate);
                    startTime.setHours(7, am * minutesPerSlot);
                    for (let k = 0; k < minutesPerSlot; k += minutesPerSlot) {
                        const timeSlot = {
                            id: (0, uuid_1.v4)(),
                            date: formattedDate,
                            time: startTime.toLocaleTimeString("en-US", {
                                hour: "numeric",
                                minute: "numeric",
                            }),
                            shift: "Morning",
                        };
                        timeSlots.push(timeSlot);
                        startTime.setMinutes(startTime.getMinutes() + minutesPerSlot);
                    }
                }
                for (let pm = 0; pm < timeSlotCountAf; pm++) {
                    const startTime = new Date(currentDate);
                    startTime.setHours(13, 30 + pm * minutesPerSlot);
                    for (let k = 0; k < minutesPerSlot; k += minutesPerSlot) {
                        const timeSlot = {
                            id: (0, uuid_1.v4)(),
                            date: formattedDate,
                            time: startTime.toLocaleTimeString("en-US", {
                                hour: "numeric",
                                minute: "numeric",
                            }),
                            shift: "Afternoon",
                        };
                        timeSlots.push(timeSlot);
                        startTime.setMinutes(startTime.getMinutes() + minutesPerSlot);
                    }
                }
                return timeSlots;
            }
            // Hàm tạo danh sách phòng và lịch trình cho mỗi phòng cố định
            function generateScheduleRoomSecond(currentDate) {
                const roomSchedules = [];
                roomLists
                    .slice(0, handleNumberRoomDef(scheduleParams.quantityWeek, numStudent))
                    .map((room) => {
                    const timeSlots = generateTimeSlotSeconds(currentDate);
                    roomSchedules.push({
                        room: room,
                        timeSlots: timeSlots,
                    });
                });
                return roomSchedules;
            }
            // Hàm tạo danh sách thời gian báo cáo cụ thể (10 people/day)
            function generateTimeSlotSeconds(currentDate) {
                const timeSlots = [];
                const minutesPerSlot = 40; //Input
                const timeSlotCountMor = 5; // Input
                const timeSlotCountAf = 5; // Input
                const formattedDate = currentDate.toLocaleDateString("en-US", {
                    weekday: "long",
                    month: "numeric",
                    day: "numeric",
                    year: "numeric",
                });
                for (let am = 0; am < timeSlotCountMor; am++) {
                    const startTime = new Date(currentDate);
                    startTime.setHours(7, am * minutesPerSlot);
                    for (let k = 0; k < minutesPerSlot; k += minutesPerSlot) {
                        const timeSlot = {
                            id: (0, uuid_1.v4)(),
                            date: formattedDate,
                            time: startTime.toLocaleTimeString("en-US", {
                                hour: "numeric",
                                minute: "numeric",
                            }),
                            shift: "Morning",
                        };
                        timeSlots.push(timeSlot);
                        startTime.setMinutes(startTime.getMinutes() + minutesPerSlot);
                    }
                }
                for (let pm = 0; pm < timeSlotCountAf; pm++) {
                    const startTime = new Date(currentDate);
                    startTime.setHours(13, 30 + pm * minutesPerSlot);
                    for (let k = 0; k < minutesPerSlot; k += minutesPerSlot) {
                        const timeSlot = {
                            id: (0, uuid_1.v4)(),
                            date: formattedDate,
                            time: startTime.toLocaleTimeString("en-US", {
                                hour: "numeric",
                                minute: "numeric",
                            }),
                            shift: "Afternoon",
                        };
                        timeSlots.push(timeSlot);
                        startTime.setMinutes(startTime.getMinutes() + minutesPerSlot);
                    }
                }
                return timeSlots;
            }
            // Hàm mục tiêu, tính toán độ chính xác dựa vào lịch trình nếu có xảy ra conflict
            function fitness(thesis, totalStudents) {
                let fitnessScore = 0;
                function findStudentConflicts(thesis) {
                    const conflicts = [];
                    const timeSlots = [];
                    for (let i = 0; i < thesis.thesis.length; i++) {
                        const timeSlotsData = thesis.thesis[i].schedule.timeSlots;
                        for (let j = 0; j < timeSlotsData.length; j++) {
                            const timeSlot = timeSlotsData[j];
                            const student = timeSlot.student;
                            if (timeSlots.includes(student) && student.id !== "") {
                                conflicts.push({
                                    student: timeSlot.student,
                                    timeSlot: timeSlot.timeSlot,
                                });
                            }
                            else {
                                timeSlots.push(student);
                            }
                        }
                    }
                    return conflicts;
                }
                const studentConflicts = findStudentConflicts(thesis);
                // console.log(studentConflicts.length);
                const numStudentConflicts = studentConflicts.length;
                const conflictRatio = numStudentConflicts / totalStudents;
                fitnessScore += 1 - conflictRatio;
                return fitnessScore;
            }
            // Lấy 70% lịch trình trong quần thể để tiến hành lai ghép và đột biến
            function selection(evaluatedPopulation) {
                const populationSize = evaluatedPopulation.length;
                const totalFitness = evaluatedPopulation.reduce((sum, individual) => sum + individual.fitness, 0);
                const selectionProbabilities = evaluatedPopulation.map((individual) => individual.fitness / totalFitness);
                const selectedPopulation = [];
                const selectedPopulationSize = Math.floor(populationSize * 0.7); // Chọn 70% số lượng cá thể
                for (let i = 0; i < selectedPopulationSize; i++) {
                    let selectedIndex = 0;
                    let randomValue = Math.random();
                    while (randomValue > 0) {
                        randomValue -= selectionProbabilities[selectedIndex];
                        selectedIndex++;
                    }
                    selectedIndex--;
                    selectedPopulation.push(evaluatedPopulation[selectedIndex]);
                }
                return selectedPopulation;
            }
            // Qúa trình lai ghép 2 cá thể trong quần thể
            function crossover(parent1, parent2) {
                // Copy information from parents
                const child1 = JSON.parse(JSON.stringify(parent1));
                const child2 = JSON.parse(JSON.stringify(parent2));
                // Select two random student positions in the child individuals
                const studentIndex1 = Math.floor(Math.random() * child1.thesis[0].schedule.timeSlots.length);
                const studentIndex2 = Math.floor(Math.random() * child2.thesis[0].schedule.timeSlots.length);
                // Check if both objects have values
                let student1 = child1.thesis[0].schedule.timeSlots[studentIndex1].student;
                let student2 = child2.thesis[0].schedule.timeSlots[studentIndex2].student;
                if (student1 && student2) {
                    const tempStudent = student1;
                    student1 = student2;
                    student2 = tempStudent;
                }
                // Return the array after swapping
                return [child1, child2];
            }
            // Qúa trình đột biến diễn ra trong quần thể với tỷ lệ 10% tức là 1-50 sv thì sẽ có đột biến
            function mutation(individual, mutationRate) {
                const timetableLength = individual.thesis.length;
                for (let i = 0; i < timetableLength; i++) {
                    if (Math.random() < mutationRate) {
                        const thesis = individual.thesis[i];
                        const timeSlotsLength = thesis.schedule.timeSlots.length;
                        if (timeSlotsLength > 0) {
                            const randomIndex = Math.floor(Math.random() * timeSlotsLength);
                            const startDate = new Date("November 27, 2023");
                            let timeSlot = thesis.schedule.timeSlots[randomIndex];
                            const currentDate = new Date(startDate);
                            const startTime = new Date(currentDate);
                            currentDate.setDate(startDate.getDate() + i);
                            const formattedDate = currentDate.toLocaleDateString("en-US", {
                                weekday: "long",
                                month: "numeric",
                                day: "numeric",
                                year: "numeric",
                            });
                            startTime.setHours(7, 0);
                            let newTimeSlot = {
                                id: (0, uuid_1.v4)(),
                                date: formattedDate,
                                time: startTime.toLocaleTimeString("en-US", {
                                    hour: "numeric",
                                    minute: "numeric",
                                }),
                                shift: "Morning",
                            };
                            timeSlot.timeSlot = newTimeSlot;
                        }
                    }
                }
                return individual;
            }
            // Chạy giải thuật Genetic algorithm
            let mutionPopulation = [];
            for (let generation = 0; generation < maxGenerations; generation++) {
                // Đánh giá độ thích nghi của quần thể
                const totalScheduleTimeSlots = initPopulation.reduce((total, individual) => {
                    const individualTimeSlots = individual.thesis.reduce((thesisTotal, thesis) => {
                        return thesisTotal + thesis.schedule.timeSlots.length;
                    }, 0);
                    return total + individualTimeSlots;
                }, 0);
                const evaluatedPopulation = initPopulation.map((individual, index) => ({
                    thesis: individual.thesis,
                    fitness: fitness(individual, totalScheduleTimeSlots),
                }));
                // Lựa chọn cá thể cho thế hệ mới
                const selectedPopulation = selection(evaluatedPopulation);
                // Tạo thế hệ mới bằng cách lai ghép và đột biến cá thể
                let newPopulation = [];
                while (newPopulation.length < populationSize) {
                    // Lựa chọn hai cá thể cha mẹ
                    const parent1 = lodash_1.default.sample(selectedPopulation);
                    const parent2 = lodash_1.default.sample(selectedPopulation);
                    // Lai ghép để tạo ra cá thể con
                    const child = crossover(parent1, parent2);
                    // Đột biến cá thể con
                    for (let i = 0; i < child.length; i++) {
                        const newChild = mutation(child[i], mutationRate);
                        newPopulation.push(newChild);
                    }
                }
                // Cập nhật quần thể mới
                mutionPopulation = newPopulation;
            }
            // Sắp xếp giảm dần theo độ chính xác
            mutionPopulation.sort((a, b) => {
                if (a.fitness !== b.fitness) {
                    return b.fitness - a.fitness; // Sắp xếp theo fitness cao nhất
                }
                else {
                    return a.thesis.length - b.thesis.length; // Nếu fitness giống nhau, sắp xếp theo độ dài bestSchedule tăng dần
                }
            });
            // Lấy ra lịch trình trong quần thể (Với độ chính xác cao nhất và độ dài thấp nhất)
            const bestSchedule = mutionPopulation[0].thesis;
            // Trả về client lịch trình tốt nhất
            // res.status(200).json(bestSchedule);
            await scheduledModel_1.ScheduleDefModel.saveScheduleDef({
                thesis: bestSchedule,
            });
        }
        catch (err) {
            console.log(err);
        }
    },
    async getSchedule(req) {
        try {
            const schedule = await scheduledModel_1.ScheduleDefModel.getSchedule();
            if (!schedule) {
                console.log("Schedule is empty");
            }
            return schedule;
        }
        catch (err) {
            console.error(err);
        }
    },
    async getCouncilInSchedule(req, res) {
        try {
            const id = req.params.id;
            const schedule = await scheduledModel_1.ScheduleDefModel.getCouncilInSchedule(id);
            if (!schedule) {
                res.status(404).json("Schedule is empty");
                return;
            }
            res.status(200).json(schedule);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getScheduleForStudent(req, res) {
        try {
            const id = req.params.id;
            const schedule = await scheduledModel_1.ScheduleDefModel.getScheduleForStudent(id);
            if (!schedule) {
                res.status(404).json("Schedule is empty");
            }
            else {
                res.status(200).json(schedule);
            }
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getScheduleForLecturer(req, res) {
        try {
            const id = req.params.id;
            const schedule = await scheduledModel_1.ScheduleDefModel.getScheduleForLecturer(id);
            if (!schedule) {
                res.status(404).json("Schedule is empty");
            }
            else {
                res.status(200).json(schedule);
            }
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
};
