package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"image_service/internal/protocol"
	"time"
)

type dbhandler struct {
	col              *mongo.Collection
	client           *mongo.Client
	checkingDuration time.Duration
}

func connect(ctx context.Context, uri, dbname, colname string) (*mongo.Client, *mongo.Collection) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	col := client.Database(dbname).Collection(colname)
	return client, col
}

type adapter struct {
	pagination int64
	uri        string
	colName    string
	dbName     string
	handler    *dbhandler
	logger     protocol.Logger
}

func (a *adapter) check(ctx context.Context) {
	tick := time.NewTicker(a.handler.checkingDuration)
	for {
		select {
		case <-tick.C:
			if a.isResetRequire(ctx) {
				a.logger.Warning("db connection issue")
				a.tryConnect(ctx)
			}
		}
	}
}

func (a *adapter) isResetRequire(ctx context.Context) bool {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	if a.handler.client == nil || a.handler.col == nil {
		return true
	}
	if err := a.handler.client.Ping(ctx, nil); err != nil {
		return true
	}
	return false
}

func (a *adapter) tryConnect(ctx context.Context) {
	client, col := connect(ctx, a.uri, a.dbName, a.colName)
	*a.handler.client = *client
	*a.handler.col = *col
}

func (a *adapter) getCol() *mongo.Collection {
	if a.handler.col != nil {
		return a.handler.col
	}
	a.tryConnect(context.TODO())
	return a.handler.col
}

func NewMongoDbAdapter(logger protocol.Logger, option ...optFunc) *adapter {
	a := adapter{
		pagination: defaultPagination,
		uri:        defaultURI,
		colName:    defaultColName,
		dbName:     defaultDbName,
		handler:    &dbhandler{checkingDuration: defaultCheckingDuration},
		logger:     logger,
	}

	for _, opt := range option {
		opt(&a)
	}
	ctx := context.Background()
	a.tryConnect(ctx)
	go a.check(ctx)
	return &a
}

type optFunc func(*adapter)

func WithURI(uri string) optFunc {
	return func(a *adapter) {
		a.uri = uri
	}
}

func WithPagination(p int64) optFunc {
	return func(a *adapter) {
		a.pagination = p
	}
}

func WithCheckingDuraion(d time.Duration) optFunc {
	return func(a *adapter) {
		a.handler.checkingDuration = d
	}
}
