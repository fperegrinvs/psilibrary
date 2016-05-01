'use strict'; // http://stackoverflow.com/questions/1335851/what-does-use-strict-do-in-javascript-and-what-is-the-reasoning-behind-it
angular.module('psilibrary.controllers')
    .controller('searchCtl', ['$scope', 'Facebook', 'searchService', function ($scope, Facebook, searchService) {
    $scope.filter = function(facet, facet_option) {
        filters = {};
        filters[facet] = ['' + facet_option];
        $scope.search(filters);
    }

    $scope.search = function(filters) {
        var query = {};

        if ($scope.query && $scope.query.length > 0){
            query['query'] = $scope.query;
        }

        if ($scope.page) {
            query['page'] = $scope.page;
        }

        if (filters) {
            $scope.filters = filters 
        }

        if ($scope.filters) {
            query['filters'] = filters;
        }        

        var call = searchService.Search(query);
        call.then(
            function(data){
                $scope.data = data;
                $scope.processNavigation(data.navigation);
            },
            function(err){
                $scope.msg = {error: err};
            });
    }

    $scope.changePage = function(page) {
        if (page < $scope.pagination.min_page || page > $scope.pagination.max_page) {
            return;
        }
        $scope.page = page;
        $scope.search();
    }

    $scope.processNavigation = function(data) {
        if (!data) {
            return;
        }

        p = {};
        p.page = data["currentPage"];
        p.total_pages = data['totalPages'];
        p.total_results = data['totalCount'];
        p.start = data['pageStart'];
        p.end = data['pageEnd'];
        p.min_page = Math.max(1, p.page - 7);
        visiblePages = p.page - p.min_page + 1;
        p.max_page = Math.min(p.total_pages, p.page + 15 - visiblePages);
        p.pages = [];

        for (var i = p.min_page; i <= p.max_page; i++) {
            pageData = {};
            pageData['name'] = i;
            pageData['current'] = p.page == i;
            p.pages.push(pageData);
        }

        $scope.pagination = p;
    }

    $scope.init = function() {
        $scope.search();
    };

    $scope.init();
}]);

