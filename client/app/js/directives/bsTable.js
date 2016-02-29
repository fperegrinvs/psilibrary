// JavaScript source code
angular.module('psilibrary.directives', [])
    .directive('bsTableControl', function () {
    return {
        restrict: 'EA',
        scope: {
            tabledata: '='
        },
        link: function (scope, element, attr) {
            var tableCreated = false;
            scope.$watch('tabledata', function (newValue, oldValue) {
                if (tableCreated && newValue === oldValue)
                    return;
                if (newValue) {
                    $(element).bootstrapTable();
                    $(element).bootstrapTable('load', scope.tabledata);
                }
                tableCreated = typeof (newValue) !== 'undefined';
            });
            $(window).resize(function () {
                if (tableCreated)
                    $(element).bootstrapTable('resetView');
            });
        }
    };
});
