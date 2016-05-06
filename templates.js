(function() {
    var module;

    try {
        // Get current templates module
        module = angular.module('templates');
    } catch (error) {
        // Or create a new one
        module = angular.module('templates', []);
    }

    module.run(["$templateCache", function($templateCache) {
        $templateCache.put('app/templates/404.html', '<h1>404</h1>');
    }]);
})();
(function() {
    var module;

    try {
        // Get current templates module
        module = angular.module('templates');
    } catch (error) {
        // Or create a new one
        module = angular.module('templates', []);
    }

    module.run(["$templateCache", function($templateCache) {
        $templateCache.put('app/templates/admin.html', '<div ng-include src=\"\'app/templates/adminHeader.html\'\"></div>\r\n\r\n<div style=\"min-height: 800px;\">\r\n<div class=\"container-fluid\" ng-show=\"!logged\">\r\n<div class=\"col-lg-12\">\r\n	<div class=\"alert alert-warning\">\r\n	  <strong>Atenção!</strong> É necessário efetuar login no facebook para continuar.\r\n	</div>\r\n\r\n    <button type=\"button\" class=\"btn btn-primary btn-large\" data-ng-show=\"!logged\" data-ng-disabled=\"!facebookReady\" data-ng-click=\"IntentLogin()\">Login with Facebook</button>\r\n</div>\r\n</div>\r\n\r\n<div class=\"container-fluid\" ng-show=\"logged && !$root.authorized\">\r\n<div class=\"col-lg-12\">\r\n  <div class=\"alert alert-danger\">\r\n    <strong>Atenção!</strong> Acesso não autorizado ao usuário {{user.name}}.\r\n  </div>\r\n\r\n  <button type=\"button\" class=\"btn btn-danger btn-large\" data-ng-show=\"logged\" data-ng-disabled=\"!facebookReady\" data-ng-click=\"logout()\">Logout</button>\r\n</div>\r\n</div>\r\n</div>');
    }]);
})();
(function() {
    var module;

    try {
        // Get current templates module
        module = angular.module('templates');
    } catch (error) {
        // Or create a new one
        module = angular.module('templates', []);
    }

    module.run(["$templateCache", function($templateCache) {
        $templateCache.put('app/templates/adminHeader.html', '<nav class=\"navbar navbar-default navbar-fixed-top\">\r\n  <div class=\"container\">\r\n    <div class=\"navbar-header\">\r\n      <button type=\"button\" class=\"navbar-toggle collapsed\" data-toggle=\"collapse\" data-target=\"#navbar\" aria-expanded=\"false\" aria-controls=\"navbar\">\r\n        <span class=\"sr-only\">Toggle navigation</span>\r\n        <span class=\"icon-bar\"></span>\r\n        <span class=\"icon-bar\"></span>\r\n        <span class=\"icon-bar\"></span>\r\n      </button>\r\n      <a class=\"navbar-brand\" href=\"\" ui-sref=\"admin\">Admin do Psi Library</a>\r\n    </div>\r\n    <div id=\"navbar\" class=\"navbar-collapse collapse\">\r\n      <ul class=\"nav navbar-nav\">\r\n        <li class=\"{{section == \'entryType\' ? \'active\' : \'\'}}\" ui-sref=\"entryType\"><a href=\"\">Tipo de Registro</a></li>\r\n        <li class=\"{{section == \'category\' ? \'active\' : \'\'}}\"><a href=\"\" ui-sref=\"category\">Categoria</a></li>\r\n        <li class=\"{{section == \'entry\' ? \'active\' : \'\'}}\"><a href=\"\" ui-sref=\"entry\">Registro</a></li>\r\n        <li class=\"{{section == \'medline\' ? \'active\' : \'\'}}\"><a href=\"\" ui-sref=\"medline\">Importação Medline</a></li>\r\n        <li><a href=\"\" ui-sref=\"home\">Voltar para o site</a></li>\r\n      </ul>\r\n      <ul class=\"nav navbar-nav navbar-right\" ng-show=\"$root.authorized\">\r\n        <li><img src=\"{{$root.user.picture.data.url}}\" /></li>\r\n        <li><a href=\"\">{{$root.user.name}}</a></li>\r\n      </ul>\r\n    </div>\r\n  </div>\r\n</nav>\r\n<br/><br/><br/>\r\n');
    }]);
})();
(function() {
    var module;

    try {
        // Get current templates module
        module = angular.module('templates');
    } catch (error) {
        // Or create a new one
        module = angular.module('templates', []);
    }

    module.run(["$templateCache", function($templateCache) {
        $templateCache.put('app/templates/categoryEdit.html', '<div ng-include src=\"\'app/templates/adminHeader.html\'\"></div>\r\n<div class=\"container-fluid\">\r\n    <div class=\"simple-form\">\r\n        <h3 ng-if=\"!data.id\" class=\"title\">Nova categoria</h3>\r\n        <h3 ng-if=\"data.id\" class=\"title\">Editando categoria</h3>\r\n        <form>\r\n            <div class=\"form-group\" ng-if=\"data.id\">\r\n                <label for=\"id\">Id</label>\r\n                <input type=\"text\" readonly=\"true\" class=\"form-control\" ng-model=\"data.id\" id=\"id\">\r\n            </div>\r\n            <div class=\"form-group\">\r\n                <label for=\"name\">Nome</label>\r\n                <input type=\"text\" class=\"form-control\" ng-model=\"data.name\" id=\"name\" placeholder=\"Nome da Categoria\">\r\n            </div>\r\n            <button type=\"submit\" class=\"btn btn-primary save\" ng-click=\"update()\">Salvar</button>\r\n        </form>\r\n    </div>\r\n</div>');
    }]);
})();
(function() {
    var module;

    try {
        // Get current templates module
        module = angular.module('templates');
    } catch (error) {
        // Or create a new one
        module = angular.module('templates', []);
    }

    module.run(["$templateCache", function($templateCache) {
        $templateCache.put('app/templates/categoryList.html', '<div ng-include src=\"\'app/templates/adminHeader.html\'\"></div>\r\n<div class=\"container-fluid\">\r\n    <h3>Categorias</h3>\r\n    <div id=\"toolbar\">\r\n        <div class=\"form-inline\" role=\"form\">\r\n            <button type=\"submit\" class=\"btn btn-primary create\" ng-click=\"create()\">Nova</button>\r\n        </div>\r\n    </div>\r\n    <table id=\"table-pagination\" data-toolbar=\"#toolbar\" data-height=\"400\" data-tabledata=\"data\" data-search=\"true\"  bs-table-control>\r\n        <thead>\r\n            <tr>\r\n                <th data-field=\"id\" data-align=\"right\" data-sortable=\"true\">ID</th>\r\n                <th data-field=\"name\" data-align=\"center\" data-sortable=\"true\">Nome</th>\r\n                <th data-field=\"id\" data-align=\"center\" data-sortable=\"true\" data-formatter=\"currentScope.formatter\">Ações</th>\r\n            </tr>\r\n        </thead>\r\n    </table>\r\n</div>');
    }]);
})();
(function() {
    var module;

    try {
        // Get current templates module
        module = angular.module('templates');
    } catch (error) {
        // Or create a new one
        module = angular.module('templates', []);
    }

    module.run(["$templateCache", function($templateCache) {
        $templateCache.put('app/templates/entryEdit.html', '<div ng-include src=\"\'app/templates/adminHeader.html\'\"></div>\r\n<div class=\"container-fluid\">\r\n<h3 ng-if=\"!data.id\" class=\"title\">Novo registro</h3>\r\n<h3 ng-if=\"data.id\" class=\"title\">Editando registro</h3>\r\n\r\n<div class=\"simple-form\">\r\n    <form name=\'dataForm\' novalidate>\r\n        <div class=\"form-group\" ng-if=\"data.id\">\r\n            <label for=\"id\">Id</label>\r\n            <input type=\"text\" readonly=\"true\" class=\"form-control\" ng-model=\"data.id\" id=\"id\" />\r\n        </div>\r\n        <div class=\"form-group\">\r\n            <label for=\"medlineId\">Id Medline</label>\r\n            <input type=\"text\" readonly=\"true\" class=\"form-control\" ng-model=\"data.medlineId\" id=\"medlineId\" name=\"medlineId\">\r\n        </div>\r\n        <div class=\"form-group\" show-errors>\r\n            <label for=\"title\">Title</label>\r\n            <input type=\"text\" class=\"form-control\" ng-model=\"data.title\" id=\"title\" name=\"title\" ng-required=\"true\">\r\n            <p class=\"help-block title-required-error\" ng-if=\"dataForm.title.$error.required\">O título é obrigatório</p>\r\n        </div>\r\n        <div class=\"form-group\" show-errors>\r\n            <label for=\"title\">Autor</label>\r\n            <input type=\"text\" class=\"form-control\" ng-model=\"data.author\" id=\"author\" name=\"author\" ng-required=\"true\">\r\n            <p class=\"help-block author-required-error\" ng-if=\"dataForm.author.$error.required\">O autor é obrigatório</p>\r\n        </div>\r\n        <div class=\"form-group\" show-errors>\r\n            <label for=\"journal\">Periódico</label>\r\n            <input type=\"text\" class=\"form-control\" ng-model=\"data.journal\" id=\"journal\" name=\"journal\">\r\n        </div>\r\n        <div class=\"form-group\" show-errors>\r\n            <label for=\"publishDate\">Data de publicação</label>\r\n            <input class=\"form-control\" ng-model=\"data.publishDateLocal\" type=\"datetime-local\" name=\"publishDate\" />\r\n        </div>\r\n        <div class=\"form-group\" show-errors>\r\n            <label for=\"entryType\">Tipo de Registro</label>\r\n            <select  name=\"entryType\" ng-options=\"entryType.name for entryType in entryTypes\" ng-model=\"data.entryType\" ng-required=\"true\"> </select>\r\n        </div>\r\n        <div class=\"form-group\">\r\n            <label for=\"categories\">Categorias</label>\r\n            <select  id=\"categories\" ng-options=\"category.name for category in categories\" multiple ng-model=\"data.categories\" bs-duallistbox non-selected-list-label=\"categorias disponíveis\" selected-list-label=\"categorias selecionadas\"> </select>\r\n        </div>\r\n        <div class=\"form-group\" show-errors>\r\n            <label for=\"abstract\">Resumo</label>\r\n            <div text-angular ng-model=\"data.abstract\" name=\"abstract\" ng-required=\"true\" />\r\n        </div>\r\n        <div class=\"form-group\">\r\n            <label for=\"content\">Conteúdo</label>\r\n            <div text-angular ng-model=\"data.content\" name=\"content\" />\r\n        </div>\r\n        <div class=\"form-group\">\r\n            <button type=\"submit\" class=\"btn btn-primary save\" ng-click=\"save()\">Salvar</button>\r\n        </div>\r\n    </form>\r\n</div>\r\n</div>');
    }]);
})();
(function() {
    var module;

    try {
        // Get current templates module
        module = angular.module('templates');
    } catch (error) {
        // Or create a new one
        module = angular.module('templates', []);
    }

    module.run(["$templateCache", function($templateCache) {
        $templateCache.put('app/templates/entryList.html', '<div ng-include src=\"\'app/templates/adminHeader.html\'\"></div>\r\n<div class=\"container-fluid\">\r\n<h3>Registros</h3>\r\n<div id=\"toolbar\">\r\n    <div class=\"form-inline\" role=\"form\">\r\n        <button type=\"submit\" class=\"btn btn-primary create\" ng-click=\"create()\">Novo</button>\r\n    </div>\r\n</div>\r\n<table id=\"table-data\" data-toolbar=\"#toolbar\" data-height=\"800\" data-tabledata=\"data\" data-search=\"true\"  bs-table-control>\r\n    <thead>\r\n        <tr>\r\n            <th data-field=\"title\" data-align=\"center\" data-sortable=\"true\" style=\"width: 10%\">Title</th>\r\n            <th data-field=\"abstract\" data-align=\"center\" data-sortable=\"true\" data-formatter=\"currentScope.abstractFormatter\">Abstract</th>\r\n            <th data-field=\"author\" data-align=\"center\" data-sortable=\"true\">Author</th>\r\n            <th data-field=\"id\" data-align=\"center\" data-sortable=\"true\" data-formatter=\"currentScope.formatter\">Ações</th>\r\n        </tr>\r\n    </thead>\r\n</table>\r\n</div>');
    }]);
})();
(function() {
    var module;

    try {
        // Get current templates module
        module = angular.module('templates');
    } catch (error) {
        // Or create a new one
        module = angular.module('templates', []);
    }

    module.run(["$templateCache", function($templateCache) {
        $templateCache.put('app/templates/entryTypeEdit.html', '<div ng-include src=\"\'app/templates/adminHeader.html\'\"></div>\r\n<div class=\"container-fluid\">\r\n<h3 ng-if=\"!data.id\">Novo tipo de registro</h3>\r\n<h3 ng-if=\"data.id\">Editando tipo de registro</h3>\r\n<div class=\"simple-form\">\r\n    <form>\r\n        <div class=\"form-group\" ng-if=\"data.id\">\r\n            <label for=\"id\">Id</label>\r\n            <input type=\"text\" readonly=\"true\" class=\"form-control\" ng-model=\"data.id\" id=\"id\">\r\n        </div>\r\n        <div class=\"form-group\">\r\n            <label for=\"name\">Nome</label>\r\n            <input type=\"text\" class=\"form-control\" ng-model=\"data.name\" id=\"name\" placeholder=\"Tipo de Registro\">\r\n        </div>\r\n        <button type=\"submit\" class=\"btn btn-primary\" ng-click=\"update()\">Salvar</button>\r\n    </form>\r\n</div>\r\n</div>');
    }]);
})();
(function() {
    var module;

    try {
        // Get current templates module
        module = angular.module('templates');
    } catch (error) {
        // Or create a new one
        module = angular.module('templates', []);
    }

    module.run(["$templateCache", function($templateCache) {
        $templateCache.put('app/templates/entryTypeList.html', '<div ng-include src=\"\'app/templates/adminHeader.html\'\"></div>\r\n<div class=\"container-fluid\">\r\n<h3>Tipos de Registro</h3>\r\n<div id=\"toolbar\">\r\n    <div class=\"form-inline\" role=\"form\">\r\n        <button type=\"submit\" class=\"btn btn-primary\" ng-click=\"create()\">Novo</button>\r\n    </div>\r\n</div>\r\n<table id=\"table-pagination\" data-toolbar=\"#toolbar\" data-height=\"400\" data-tabledata=\"data\" data-search=\"true\"  bs-table-control>\r\n    <thead>\r\n        <tr>\r\n            <th data-field=\"id\" data-align=\"right\" data-sortable=\"true\">ID</th>\r\n            <th data-field=\"name\" data-align=\"center\" data-sortable=\"true\">Nome</th>\r\n            <th data-field=\"id\" data-align=\"center\" data-sortable=\"true\" data-formatter=\"currentScope.formatter\">Ações</th>\r\n        </tr>\r\n    </thead>\r\n</table>\r\n</div>');
    }]);
})();
(function() {
    var module;

    try {
        // Get current templates module
        module = angular.module('templates');
    } catch (error) {
        // Or create a new one
        module = angular.module('templates', []);
    }

    module.run(["$templateCache", function($templateCache) {
        $templateCache.put('app/templates/home.html', '<h1>Sample home template</h1>');
    }]);
})();
(function() {
    var module;

    try {
        // Get current templates module
        module = angular.module('templates');
    } catch (error) {
        // Or create a new one
        module = angular.module('templates', []);
    }

    module.run(["$templateCache", function($templateCache) {
        $templateCache.put('app/templates/medline.html', '<div ng-include src=\"\'app/templates/adminHeader.html\'\"></div>\r\n<div class=\"simple-form\">\r\n	<h3>Importar registros do Medline</h3>\r\n    <form>\r\n        <div class=\"form-group\">\r\n            <label for=\"name\">Nome</label>\r\n            <textarea type=\"text\" class=\"form-control\" ng-model=\"xml\" id=\"xml\" placeholder=\"XML\" style=\"height: 500px;\" />\r\n        </div>\r\n        <button type=\"submit\" class=\"btn btn-primary\" ng-click=\"import()\">Importar</button>\r\n    </form>\r\n</div>\r\n');
    }]);
})();
(function() {
    var module;

    try {
        // Get current templates module
        module = angular.module('templates');
    } catch (error) {
        // Or create a new one
        module = angular.module('templates', []);
    }

    module.run(["$templateCache", function($templateCache) {
        $templateCache.put('app/templates/search.html', '<div ng-include src=\"\'app/templates/siteHeader.html\'\"></div>\r\n\r\n<div class=\"container-fluid\">\r\n\r\n<div class=\"col-lg-12\">\r\n<div id=\"search-bar\">\r\n    <div class=\"input-group\">\r\n        <input type=\"text\" name=\"query\" id=\"query\" ng-model=\"query\" class=\"search-query form-control\" placeholder=\"Buscar\" />\r\n        <span class=\"input-group-btn\">\r\n            <button class=\"btn btn-danger\" type=\"button\" ng-click=\"searchClick()\">\r\n                <span class=\" glyphicon glyphicon-search\"></span>\r\n            </button>\r\n        </span>\r\n    </div>\r\n</div>\r\n</div>\r\n\r\n<!-- left -->\r\n<div class=\"col-lg-3\">\r\n	<div class=\"panel panel-primary\" ng-repeat=\"facet in data.facets\">\r\n		<div class=\"panel-heading\">\r\n			{{facet.name}}\r\n		</div>\r\n		<div class=\"panel-body\">\r\n		<div ng-repeat=\"option in facet.options\" ng-class=\"\'facet-\' + facet.id + \'-item\'\" class=\"btn btn-default btn-block {{option.isSelected ? \'active\' : \'\'}}\" ng-click=\"filter(facet[\'id\'], option[\'id\'], option[\'name\'])\">{{option.name}}</div>\r\n		</div>\r\n	</div>\r\n</div>\r\n\r\n<!-- right -->\r\n<div class=\"col-lg-9\">\r\n	<div class=\"\">\r\n		<div ng-repeat=\"result in data.results\" class=\"panel panel-primary result-item\">\r\n			<div class=\"result-title panel-heading\">\r\n				{{result.title}}\r\n			</div>\r\n	        <div class=\"panel-body\">\r\n				<h3 class=\"result-author\">\r\n					{{result.author}}\r\n				</h3>\r\n				<h4 class=\"result-journal\">\r\n					{{result.journal}}\r\n				</h4>\r\n				<div class=\"result-abstract text-justify\" ng-bind-html=\"result.abstract_safe\" />\r\n				<div>\r\n				<br/>\r\n				<button type=\"button\" ui-sref=\"view({id: result.id})\" class=\"btn btn-primary\">Ler Mais</button>\r\n				</div>\r\n			</div>\r\n		</div>\r\n	</div>\r\n	<div id=\"pagination\" class=\"col-lg-12 text-center\">\r\n		<nav>\r\n		  <ul class=\"pagination\">\r\n		    <li class=\"page-item\">\r\n		      <span class=\"page-link previous-page\" href=\"#\" aria-label=\"Anterior\">\r\n		        <span aria-hidden=\"true\">&laquo;</span>\r\n		        <span class=\"sr-only\" ng-click=\"changePage(page + 1)\">Anterior</span>\r\n		      </span>\r\n		    </li>\r\n		    \r\n		    <li ng-repeat=\"page in pagination.pages\" class=\"page-item\" \r\n		    	ng-class=\"{\'page-item\': true, \'active\': page.current}\">\r\n		    	<span class=\"page-link\" href=\"#\" ng-click=\"changePage(page.name)\">{{page.name}}</span>\r\n		    </li>\r\n\r\n		    <li class=\"page-item\">\r\n		      <span class=\"page-link next-page\" href=\"#\" aria-label=\"Next\" ng-click=\"changePage(page + 1)\">\r\n		        <span aria-hidden=\"true\">&raquo;</span>\r\n		        <span class=\"sr-only\">Next</span>\r\n		      </span>\r\n		    </li>\r\n		  </ul>\r\n		</nav>\r\n	</div>\r\n</div>\r\n\r\n</div>');
    }]);
})();
(function() {
    var module;

    try {
        // Get current templates module
        module = angular.module('templates');
    } catch (error) {
        // Or create a new one
        module = angular.module('templates', []);
    }

    module.run(["$templateCache", function($templateCache) {
        $templateCache.put('app/templates/siteHeader.html', '<div class=\"container-fluid\">\r\n<div class=\"col-lg-12\">\r\n	<h1>Psi Library</h1>\r\n</div>\r\n<div class=\"col-lg-12\">\r\n<ol class=\"breadcrumb\" style=\"margin-bottom: 5px;\">\r\n  <li ng-show=\"current == \'Home\'\" class=\"active\">Home</li>\r\n  <li ng-show=\"current != \'Home\'\" class=\"active\"><a href=\"/\">Home</a></li>\r\n  <li ng-show=\"current != \'Home\'\" class=\"active\">{{current}}</li>\r\n</ol>\r\n</div>\r\n</div>\r\n<br/>\r\n');
    }]);
})();
(function() {
    var module;

    try {
        // Get current templates module
        module = angular.module('templates');
    } catch (error) {
        // Or create a new one
        module = angular.module('templates', []);
    }

    module.run(["$templateCache", function($templateCache) {
        $templateCache.put('app/templates/view.html', '<div ng-include src=\"\'app/templates/siteHeader.html\'\"></div>\r\n\r\n<div class=\"container-fluid\">\r\n<div class=\"col-lg-12\">\r\n	<div class=\"panel panel-primary result-item\">\r\n		<div class=\"data-title panel-heading\">\r\n			{{data.title}}\r\n		</div>\r\n        <div class=\"panel-body\">\r\n			<h3 class=\"data-author\">\r\n				{{data.author}}\r\n			</h3>\r\n			<h4 class=\"result-journal\">\r\n				{{data.journal}}\r\n			</h4>\r\n			<div class=\"result-abstract text-justify\" ng-bind-html=\"data.abstract_safe\" />\r\n			<div class=\"result-content text-justify\" ng-bind-html=\"data.content_safe\" />\r\n			<div>\r\n			<br/>\r\n			<button type=\"button\" ui-sref=\"home\" class=\"btn btn-primary\">Voltar</button>\r\n			</div>\r\n		</div>\r\n	</div>\r\n</div>\r\n</div>');
    }]);
})();

//# sourceMappingURL=templates.js.map