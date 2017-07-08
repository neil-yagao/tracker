var WorkoutTemplate = (function () {
    function WorkoutTemplate() {
        this.movements = [];
        this.startAt = new Date().toString();
    }
    WorkoutTemplate.prototype.addExercise = function (active) {
        this.movements.push(active);
    };
    return WorkoutTemplate;
}());
export { WorkoutTemplate };
//# sourceMappingURL=workout-template.js.map